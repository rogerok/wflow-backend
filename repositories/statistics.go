package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
	"log"
	"time"
)

type StatisticsRepository interface {
	GetUserStatistics(userId uuid.UUID) (*models.UserStatistics, error)
	GetGoalStatistics(goalId string) (*models.GoalStatistics, error)
	GetFullProfileChartData(userId uuid.UUID) (*models.FullProfileChartData, error)
	GetGoalsChart(userID uuid.UUID) ([]models.GoalsChart, error)
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
     ),
     goal_report_days AS (
         SELECT
             g.id as goal_id,
             COUNT(DISTINCT DATE(r.created_at)) + 1 AS report_days
         FROM goals g
                  LEFT JOIN reports r ON g.id = r.goal_id
         WHERE g.is_finished = true
         GROUP BY g.id
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
        WHEN (SELECT COUNT(*) FROM goal_report_days) = 0 THEN 0
        ELSE (SELECT AVG(report_days) FROM goal_report_days)
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
    COUNT(DISTINCT d.report_date) as total_days_with_activity,
    CASE
        WHEN COUNT(DISTINCT g.id) = 0 THEN 0
        ELSE (COUNT(DISTINCT g.id) FILTER (WHERE g.is_finished = true) * 100.0 / COUNT(DISTINCT g.id))
        END as goal_completion_rate,

	CASE
		WHEN COUNT(DISTINCT d.report_date) = 0 THEN 0
		WHEN MIN(r.created_at::date) IS NULL THEN 0
		ELSE (LEAST(COUNT(DISTINCT d.report_date), (CURRENT_DATE - MIN(r.created_at::date))::int) * 100.0 /
			  NULLIF((CURRENT_DATE - MIN(r.created_at::date))::int, 0))
	END as activity_consistency_rate,

    CASE
        WHEN SUM(g.goal_words) = 0 OR SUM(g.goal_words) IS NULL THEN 0
        ELSE (SUM(g.written_words) * 100.0 / SUM(g.goal_words))
        END as overall_goal_progress_rate,

    CASE
        WHEN COUNT(DISTINCT g.id) FILTER (WHERE g.is_expired = true) = 0 THEN 0
        ELSE (COUNT(DISTINCT g.id) FILTER (WHERE g.is_expired = true AND g.is_finished = true) * 100.0 /
              NULLIF(COUNT(DISTINCT g.id) FILTER (WHERE g.is_expired = true), 0))
        END as expired_goals_completion_rate,

    CASE
        WHEN COUNT(DISTINCT g.id) = 0 THEN 0
        ELSE (COUNT(DISTINCT g.id) FILTER (WHERE g.written_words >= g.goal_words) * 100.0 /
              COUNT(DISTINCT g.id))
        END as overachievement_rate,

    CASE
        WHEN COUNT(DISTINCT r.id) = 0 THEN 0
        ELSE (COUNT(DISTINCT r.id) FILTER (WHERE r.words_amount >
                                                 (SELECT AVG(words_amount) FROM reports WHERE user_id = $1)) * 100.0 /
              COUNT(DISTINCT r.id))
        END as above_average_reports_rate
FROM users u
         LEFT JOIN books b ON u.id = b.user_id
         LEFT JOIN goals g ON b.id = g.book_id
         LEFT JOIN reports r ON g.id = r.goal_id
         LEFT JOIN daily_words d ON DATE(r.created_at) = d.report_date
WHERE u.id = $1
GROUP BY u.id;
`

	err := r.db.Get(statistic, query, userId)

	if err != nil {
		return nil, err
	}

	return statistic, nil

}

func (r *statisticsRepository) GetGoalStatistics(goalId string) (*models.GoalStatistics, error) {
	statistic := &models.GoalStatistics{}

	query := `WITH daily_reports AS (
    SELECT
        DATE(created_at) as report_date,
        SUM(words_amount) as daily_words
    FROM reports
    WHERE goal_id = $1
    GROUP BY DATE(created_at)
),
     goal_progress AS (
         SELECT
             g.id as goal_id,
             g.book_id,
             g.goal_words,
             g.start_date,
             g.end_date,
             -- Общие слова, написанные по всем отчетам
             COALESCE(SUM(r.words_amount), 0) as total_words_written,
             -- Дата первого отчета
             MIN(r.created_at) as first_report_date,
             -- Количество дней с первого отчета по сегодняшний день
             GREATEST(1, EXTRACT(DAY FROM CURRENT_DATE - MIN(r.created_at))) as days_active,
             -- Среднее количество слов в день с первого отчета
             COALESCE(SUM(r.words_amount), 0) / GREATEST(1, EXTRACT(DAY FROM CURRENT_DATE - MIN(r.created_at))) as average_words_per_day
         FROM goals g
                  LEFT JOIN reports r ON g.id = r.goal_id
         WHERE g.id = $1
         GROUP BY g.id, g.book_id, g.goal_words, g.start_date, g.end_date
     )
SELECT
    g.id as goal_id,
    g.book_id,
    gp.total_words_written,
    CASE
        WHEN g.goal_words = 0 THEN 0
        ELSE (gp.total_words_written / g.goal_words) * 100
        END as percentage_complete,
    GREATEST(g.goal_words - gp.total_words_written, 0) as remaining_words,
    CASE
        WHEN CURRENT_DATE > g.end_date THEN 0
        WHEN g.is_expired THEN 0
        WHEN g.is_finished THEN 0
        WHEN g.goal_words <= gp.total_words_written THEN 0
        ELSE (g.goal_words - gp.total_words_written) /
             GREATEST(1, EXTRACT(EPOCH FROM (g.end_date - CURRENT_DATE)) / 86400)
        END as daily_words_required,
    CASE
        WHEN CURRENT_DATE > g.end_date THEN 0
        WHEN g.is_expired THEN 0
        WHEN g.is_finished THEN 0
        ELSE GREATEST(0, (CURRENT_DATE - g.start_date::DATE) + 1)::int
        END as days_elapsed,
    -- Количество отчетов
    COUNT(DISTINCT r.id) as reports_count,
    -- Среднее количество слов на отчет
    CASE
        WHEN COUNT(DISTINCT r.id) = 0 THEN 0
        ELSE SUM(r.words_amount) / COUNT(DISTINCT r.id)
        END as average_words_per_report,
    GREATEST(0, EXTRACT(DAY FROM (g.end_date - CURRENT_DATE)) + 1)::int as days_remaining,
    -- Среднее количество слов в день на основе активности пользователя с даты первого отчета
    gp.average_words_per_day as average_words_per_day,
    -- Примерная дата окончания цели, рассчитанная на основе реальной активности
    CASE
        WHEN gp.total_words_written >= g.goal_words THEN
            g.end_date -- Если цель выполнена, то дата окончания остается прежней
        ELSE
                    CURRENT_DATE + INTERVAL '1 day' *
                                   GREATEST(0, (g.goal_words - gp.total_words_written) / NULLIF(gp.average_words_per_day, 0))
        END as estimated_end_date
FROM goals g
         LEFT JOIN goal_progress gp ON g.id = gp.goal_id
         LEFT JOIN reports r ON g.id = r.goal_id
         LEFT JOIN daily_reports dr ON DATE(r.created_at) = dr.report_date
WHERE g.id = $1
GROUP BY g.id, g.book_id, g.goal_words, g.start_date, g.end_date, g.words_per_day, gp.total_words_written, gp.average_words_per_day
`
	err := r.db.Get(statistic, query, goalId)

	if err != nil {
		return nil, err
	}
	return statistic, err
}

func (r *statisticsRepository) getDateRangeForPeriod(period string) (time.Time, time.Time) {
	now := time.Now()
	endDate := now

	switch period {
	case "week":
		startDate := now.AddDate(0, 0, -7)
		return startDate, endDate
	case "month":
		startDate := now.AddDate(0, -1, 0)
		return startDate, endDate
	case "quarter":
		startDate := now.AddDate(0, -3, 0)
		return startDate, endDate
	case "year":
		startDate := now.AddDate(-1, 0, 0)
		return startDate, endDate
	case "all_time":
		// Get the user's first report date or default to 1 year ago
		var firstReportDate time.Time
		query := `
		SELECT MIN(created_at) AS first_date
		FROM reports
		WHERE user_id = $1
		`
		err := r.db.GetContext(context.Background(), &firstReportDate, query, period)
		if err != nil || firstReportDate.IsZero() {
			firstReportDate = now.AddDate(-1, 0, 0)
		}
		return firstReportDate, endDate
	default: // Default to last 30 days
		startDate := now.AddDate(0, 0, -30)
		return startDate, endDate
	}
}

func (r *statisticsRepository) GetGoalsChart(userID uuid.UUID) ([]models.GoalsChart, error) {
	var goals []models.GoalsChart
	query := `
WITH daily_reports AS (
    SELECT
        DATE(created_at) as report_date,
        SUM(words_amount) as daily_words
    FROM reports
    GROUP BY DATE(created_at)
)
SELECT
    g.id as goal_id,
    g.book_id,
    g.created_at,
    g.is_finished,
    g.is_expired,
    g.title as goal_title,
    COALESCE(SUM(r.words_amount), 0) as total_words_written,
    CASE
        WHEN g.goal_words = 0 THEN 0
        ELSE (COALESCE(SUM(r.words_amount), 0) / g.goal_words) * 100
        END as percentage_complete,
    GREATEST(g.goal_words - COALESCE(SUM(r.words_amount), 0), 0) as remaining_words,
    CASE
        WHEN CURRENT_DATE > g.end_date THEN 0
        WHEN g.is_expired THEN 0
        WHEN g.is_finished THEN 0
        WHEN g.goal_words <= COALESCE(SUM(r.words_amount), 0) THEN 0
        ELSE (g.goal_words - COALESCE(SUM(r.words_amount), 0)) /
             GREATEST(1, EXTRACT(EPOCH FROM (g.end_date - CURRENT_DATE)) / 86400)
        END as daily_words_required,
    CASE
        WHEN CURRENT_DATE > g.end_date THEN 0
        WHEN g.is_expired THEN 0
        WHEN g.is_finished THEN 0
        ELSE GREATEST(0, (CURRENT_DATE - g.start_date::DATE) + 1)::int
        END as days_elapsed,
    GREATEST(0, EXTRACT(DAY FROM (g.end_date - CURRENT_DATE)) + 1)::int as days_remaining,
    CASE
        WHEN COUNT(DISTINCT dr.report_date) = 0 THEN 0
        ELSE COALESCE(SUM(r.words_amount), 0) / COUNT(DISTINCT dr.report_date)
        END as average_words_per_day,
    COUNT(DISTINCT r.id) as reports_count,
    CASE
        WHEN g.words_per_day = 0 THEN 0
        WHEN g.written_words = 0 THEN 0
        ELSE (
                 CASE
                     WHEN COUNT(DISTINCT dr.report_date) = 0 THEN 0
                     ELSE COALESCE(SUM(r.words_amount), 0) / COUNT(DISTINCT dr.report_date)
                     END
                 ) / g.words_per_day * 100 - 100
        END as trend_compared_to_target,
    b.book_name AS book_title
FROM goals g
         LEFT JOIN reports r ON g.id = r.goal_id
         LEFT JOIN daily_reports dr ON DATE(r.created_at) = dr.report_date
         LEFT JOIN books b ON g.book_id = b.id
WHERE g.user_id = $1
GROUP BY b.book_name, g.id, g.book_id, g.goal_words, g.start_date, g.end_date, g.words_per_day, g.created_at, g.title
ORDER BY g.created_at;
    `

	err := r.db.Select(&goals, query, userID)
	if err != nil {
		log.Printf("Error fetching goals for user %s: %v", userID, err)
		return nil, err
	}
	return goals, nil
}

func (r *statisticsRepository) getReportsForGoal(goalId uuid.UUID) ([]models.ReportStatistic, error) {
	var reports []models.ReportStatistic

	query := `
		SELECT r.id, r.goal_id, r.created_at, r.updated_at, r.words_amount, r.book_id, b.book_name, g.goal_words, g.title as goal_title
		FROM reports r
				 LEFT JOIN goals g ON g.id = r.goal_id
				 LEFT JOIN books b ON b.id = r.book_id
		WHERE goal_id = $1
		GROUP BY b.book_name, r.id, r.book_id,r.created_at, r.updated_at, r.words_amount, g.goal_words, g.title
		ORDER BY r.created_at;
	`

	err := r.db.Select(&reports, query, goalId)
	if err != nil {
		log.Printf("Error fetching reports for goal %s: %v", goalId, err)

		return nil, err
	}

	return reports, nil
}

func (r *statisticsRepository) calculateCumulativeProgress(goalId uuid.UUID) ([]models.ProgressPoint, error) {
	// Fetch all reports for the goal
	reports, err := r.getReportsForGoal(goalId)
	if err != nil {
		return nil, err
	}

	var cumulativeProgress []models.ProgressPoint
	var totalWordsWritten float64
	totalGoalWords := reports[0].GoalWords

	for _, report := range reports {
		totalWordsWritten += report.WordsAmount
		completionPercent := (totalWordsWritten / totalGoalWords) * 100

		cumulativeProgress = append(cumulativeProgress, models.ProgressPoint{
			Date:              report.CreatedAt,
			TotalWords:        totalWordsWritten,
			TargetTotalWords:  totalGoalWords,
			CompletionPercent: completionPercent,
			GoalTitle:         report.GoalTitle,
			BookName:          report.BookName,
			GoalId:            report.GoalId,
			BookId:            report.BookId,
		})
	}

	return cumulativeProgress, nil
}

func (r *statisticsRepository) GetFullProfileChartData(userId uuid.UUID) (*models.FullProfileChartData, error) {
	goals, err := r.GetGoalsChart(userId)

	if err != nil {
		return nil, err
	}
	var cumulativeProgress []models.ProgressPoint

	for _, goal := range goals {
		cumulative, err := r.calculateCumulativeProgress(goal.GoalID)
		if err != nil {
			return nil, err
		}
		cumulativeProgress = append(cumulativeProgress, cumulative...)

	}

	return &models.FullProfileChartData{
		CumulativeProgress: cumulativeProgress,
		GoalCompletion:     goals,
	}, nil
}
