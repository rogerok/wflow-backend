package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type GoalsRepository interface {
	Create(goal *models.Goals) (id *string, err error)
	Edit(goal *forms.GoalEditForm) (goalStats *models.GoalUpdateResponse, err error)
	Delete(goalId string, userId string) (status bool, err error)
	GetById(id string) (goal *models.Goals, err error)
	GetList(params *models.GoalsQueryParams) (goals *[]models.Goals, err error)
	RecalculateGoal(wordsAmount float64, goalId string) (goalStats *models.GoalStats, err error)
	RecalculateGoals()
}

type goalsRepository struct {
	db *sqlx.DB
}

func NewGoalsRepository(db *sqlx.DB) GoalsRepository { return &goalsRepository{db: db} }

func (r *goalsRepository) Create(goal *models.Goals) (id *string, err error) {

	query := `INSERT INTO goals 
    						(book_id, end_date, goal_words, start_date, title, user_id, description, words_per_day, written_words)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = r.db.QueryRow(query, goal.BookId, goal.EndDate, goal.GoalWords, goal.StartDate, goal.Title, goal.UserId, goal.Description, goal.WordsPerDay, goal.WrittenWords).Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *goalsRepository) Edit(goal *forms.GoalEditForm) (goalStats *models.GoalUpdateResponse, err error) {
	goalStats = &models.GoalUpdateResponse{}

	query := `
			WITH calculated AS (
				SELECT id,
					   COALESCE(($1::INTEGER - written_words) / NULLIF(ABS(EXTRACT(DAY FROM ($2::TIMESTAMP - $3::TIMESTAMP))) + 1, 0), 0)
					   AS calculated_words_per_day
				FROM goals
				WHERE id = $4
			)
			UPDATE goals
			SET goal_words = $1,
				end_date = $2,
				start_date = $3,
				description = $5,
				title = $6,
				words_per_day = calculated_words_per_day
			FROM calculated
			WHERE goals.id = calculated.id 
			AND user_id = $7 
			RETURNING words_per_day, goal_words;
`
	rows, err := r.db.Queryx(query, goal.GoalWords, goal.EndDate, goal.StartDate, goal.GoalId, goal.Description, goal.Title, goal.UserId)

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&goalStats.WordsPerDay, &goalStats.GoalWords)
			if err != nil {
				return nil, err
			}
		}
	}

	return goalStats, err

}

func (r *goalsRepository) Delete(goalId string, userId string) (status bool, err error) {

	query := `DELETE FROM goals WHERE id = $1 AND user_id = $2`

	_, err = r.db.Exec(query, goalId, userId)

	if err != nil {
		fmt.Printf("Error deleting goal %v. %v", goalId, err.Error())
		return false, err

	}

	return true, nil
}

func (r *goalsRepository) GetById(id string) (goal *models.Goals, err error) {
	query := `SELECT created_at, updated_at, end_date, goal_words, id, book_id, is_finished, start_date, title, description, written_words, words_per_day, is_expired FROM goals WHERE id=$1`

	goal = &models.Goals{}

	err = r.db.Get(goal, query, id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}

func (r *goalsRepository) GetList(params *models.GoalsQueryParams) (goals *[]models.Goals, err error) {

	query := `
				SELECT
					created_at, updated_at, end_date, goal_words, id,
					is_finished, start_date, title, description, written_words,
					words_per_day, is_expired, book_id
				FROM goals WHERE user_id=$1
			`

	var queryParams []interface{}
	queryParams = append(queryParams, params.UserId)

	if params.BookId != nil && *params.BookId != "" {
		query += " AND book_id=$2"
		queryParams = append(queryParams, *params.BookId)
	}

	query += utils.GetAllowedOrderBy(params.OrderBy)

	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)

	goals = &[]models.Goals{}

	if !selectAll {
		query += utils.GetOffsetLimitQuery(params.PerPage, offset)
	}

	err = r.db.Select(goals, query, queryParams...)

	if err != nil {
		return nil, err
	}

	return goals, nil

}

func (r *goalsRepository) RecalculateGoal(wordsAmount float64, goalId string) (goalStats *models.GoalStats, err error) {

	goalStats = &models.GoalStats{}

	query := `
				
	WITH calculated AS (
					SELECT id,
						   written_words + $1 AS updated_written_words,
						   EXTRACT(DAY FROM end_date - start_date) AS total_days,
						   (goal_words - (written_words + $1)) / NULLIF(EXTRACT(DAY FROM end_date - start_date), 0) AS calculated_words_per_day
					FROM goals
					WHERE id = $2
				)
				UPDATE goals
				SET written_words = updated_written_words,
				    is_finished = written_words + $1 >= goal_words,
					words_per_day = 
						CASE
							WHEN calculated_words_per_day < 1 THEN 0
							WHEN calculated_words_per_day * updated_written_words < goal_words
								THEN calculated_words_per_day + 0.1
							ELSE calculated_words_per_day
						END
				FROM calculated
				WHERE goals.id = calculated.id RETURNING written_words, words_per_day`

	rows, err := r.db.Queryx(query, wordsAmount, goalId)

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&goalStats.WrittenWords, &goalStats.WordsPerDay)
			if err != nil {
				return nil, err
			}
		}
	}

	return goalStats, err
}

func (r *goalsRepository) RecalculateGoals() {
	query := `
		WITH calculated AS (
			SELECT
				id,
				goal_words - written_words AS written_goal_difference,
				CASE
					WHEN EXTRACT(DAY FROM (end_date - now() + INTERVAL '1 day')) < 1 THEN NULL
					ELSE (goal_words - written_words) / EXTRACT(DAY FROM (end_date - now() + INTERVAL '1 day'))
					END AS calculated_words_per_day
			FROM goals
		)
		UPDATE goals
		SET
			words_per_day =
				CASE
				    WHEN EXTRACT(DAY FROM (end_date - now() + INTERVAL '1 day')) < 0 THEN 0
				    WHEN calculated.written_goal_difference < 0 THEN 0
					WHEN calculated.written_goal_difference < 1 THEN words_per_day
					WHEN calculated.calculated_words_per_day IS NULL THEN words_per_day
					ELSE calculated.calculated_words_per_day
					END,
			is_expired =
				CASE
					WHEN EXTRACT(DAY FROM (end_date - now() + INTERVAL '1 day')) < 0 AND calculated.written_goal_difference > 0 THEN true
					ELSE false
					END,
			is_finished =
				CASE
					WHEN goals.written_words >= goals.goal_words THEN true
					ELSE false
					END
		FROM calculated
		WHERE goals.id = calculated.id;
			`

	_, err := r.db.Queryx(query)

	if err != nil {
		fmt.Println(err)
	}
}
