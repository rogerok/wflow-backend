package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type GoalsRepository interface {
	Create(goal *models.Goals) (id *string, err error)
	GetById(id string) (goal *models.Goals, err error)
	GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.Goals, err error)
	RecalculateGoal(wordsAmount int, goalId string) (err error)
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

func (r *goalsRepository) GetById(id string) (goal *models.Goals, err error) {
	query := `SELECT created_at, updated_at, end_date, goal_words, id, book_id, is_finished, start_date, title, description, written_words, words_per_day, is_expired FROM goals WHERE id=$1`

	goal = &models.Goals{}

	err = r.db.Get(goal, query, id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}

func (r *goalsRepository) GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.Goals, err error) {

	query := `SELECT created_at, updated_at, end_date, goal_words, id, is_finished, start_date, title, description, written_words, words_per_day, is_expired FROM goals WHERE book_id=$1`
	query += utils.GetAllowedOrderBy(params.OrderBy)

	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)

	goals = &[]models.Goals{}

	if selectAll {
		err = r.db.Select(goals, query, params.BookId)
	} else {
		query += utils.GetOffsetLimitQuery(params.PerPage, offset)
		err = r.db.Select(goals, query, params.BookId)
	}

	if err != nil {
		return nil, err
	}

	return goals, nil

}

func (r *goalsRepository) RecalculateGoal(wordsAmount int, goalId string) (err error) {

	query := `
				WITH calculated AS (SELECT id,
										   written_words + $1                                   AS updated_written_words,
										   ROUND((goal_words / (written_words + $1)) * 10) / 10 AS calculated_words_per_day
									FROM goals
									WHERE id = $2)
				UPDATE goals
				SET written_words = updated_written_words,
					words_per_day =
						CASE
							WHEN calculated_words_per_day < 1 THEN 0
							WHEN calculated_words_per_day * updated_written_words < goal_words
								THEN calculated_words_per_day + 0.1
							ELSE calculated_words_per_day
							END
				FROM calculated
				WHERE goals.id = calculated.id;
				`

	_, err = r.db.Exec(query, wordsAmount, goalId)

	return err
}
