package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type StatisticsRepository interface {
	GetUserStatistics(userId uuid.UUID) (*models.UserStatistics, error)
	//GetGoalStatistics(ctx context.Context, goalID uuid.UUID) (*models.GoalStatistics, error)
	//GetUserDailyWordCounts(ctx context.Context, userID uuid.UUID, startDate, endDate time.Time) ([]models.DayStats, error)
}

type statisticsRepository struct {
	db *sqlx.DB
}

func NewStatisticsRepository(db *sqlx.DB) StatisticsRepository {
	return &statisticsRepository{db: db}
}

func (r *statisticsRepository) GetUserStatistics(userId uuid.UUID) (*models.UserStatistics, error) {
	statistic := &models.UserStatistics{}

	query := `WITH daily_words AS (
	    SELECT
	        DATE(created_at) as report_date,
	        SUM(words_amount) as daily_words
	    FROM reports
	    WHERE user_id = $1
	    GROUP BY DATE(created_at)
	),
	streaks AS (
	    SELECT
	        report_date,
	        report_date - (ROW_NUMBER() OVER (ORDER BY report_date) * INTERVAL '1 day') AS streak_group
	    FROM daily_words
	),
	streak_lengths AS (
	    SELECT
	        streak_group,
	        COUNT(*) as streak_length
	    FROM streaks
	    GROUP BY streak_group
	)
	SELECT
	    $1 as user_id,
	    COALESCE(SUM(r.words_amount), 0) as total_words,
	    COUNT(DISTINCT b.id) as total_books,
	    COUNT(DISTINCT g.id) as total_goals,
	    COUNT(DISTINCT g.id) FILTER (WHERE g.is_finished = true) as completed_goals,
	    COUNT(DISTINCT r.id) as total_reports,
	    CASE
	        WHEN COUNT(DISTINCT d.report_date) = 0 THEN 0
	        ELSE COALESCE(SUM(r.words_amount), 0) / COUNT(DISTINCT d.report_date)
	    END as average_words_per_day,
	    CASE
	        WHEN COUNT(DISTINCT r.id) = 0 THEN 0
	        ELSE COALESCE(SUM(r.words_amount), 0) / COUNT(DISTINCT r.id)
	    END as average_words_per_report,
	    CASE
	        WHEN COUNT(DISTINCT g.id) FILTER (WHERE g.is_finished = true) = 0 THEN 0
	        ELSE AVG(EXTRACT(EPOCH FROM (g.end_date - g.start_date)) / 86400) FILTER (WHERE g.is_finished = true)
	    END as average_days_to_complete,
	    (SELECT report_date FROM daily_words ORDER BY daily_words DESC LIMIT 1) as most_productive_day,
	    COALESCE((SELECT MAX(daily_words) FROM daily_words), 0) as max_words_in_day,
	    COALESCE((
	        SELECT MAX(streak_length)
	        FROM streak_lengths
	        WHERE streak_group = (
	            SELECT streak_group
	            FROM streaks
	            WHERE report_date = CURRENT_DATE
	            LIMIT 1
	        )
	    ), 0) as current_streak,
	    COALESCE((SELECT MAX(streak_length) FROM streak_lengths), 0) as longest_streak,
	    COUNT(DISTINCT d.report_date) as total_days_with_activity
	FROM users u
	LEFT JOIN books b ON u.id = b.user_id
	LEFT JOIN goals g ON b.id = g.book_id
	LEFT JOIN reports r ON g.id = r.goal_id
	LEFT JOIN daily_words d ON DATE(r.created_at) = d.report_date  -- ✅ Fix applied here
	WHERE u.id = $1
	GROUP BY u.id;
	`

	err := r.db.Get(statistic, query, userId)

	if err != nil {
		return nil, err
	}

	return statistic, nil

}
