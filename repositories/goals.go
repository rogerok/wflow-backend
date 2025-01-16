package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type GoalsRepository interface {
	Create(goal *models.GoalsModel) (id *string, err error)
	GetById(id string) (goal *models.GoalsModel, err error)
	GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.GoalsModel, err error)
}

type goalsRepository struct {
	db *sqlx.DB
}

func NewGoalsRepository(db *sqlx.DB) GoalsRepository { return &goalsRepository{db: db} }

func (r *goalsRepository) Create(goal *models.GoalsModel) (id *string, err error) {

	query := `INSERT INTO goals 
    						(book_id, end_date, goal_words, start_date, title, user_id, description, words_per_day, written_words)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = r.db.QueryRow(query, goal.BookId, goal.EndDate, goal.GoalWords, goal.StartDate, goal.Title, goal.UserId, goal.Description, goal.WordsPerDay, goal.WrittenWords).Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *goalsRepository) GetById(id string) (goal *models.GoalsModel, err error) {
	query := `SELECT created_at, updated_at, end_date, goal_words, id, book_id, is_finished, start_date, title, description, written_words, words_per_day, is_expired FROM goals WHERE id=$1`

	goal = &models.GoalsModel{}

	err = r.db.Get(goal, query, id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}

func (r *goalsRepository) GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.GoalsModel, err error) {

	query := `SELECT created_at, updated_at, end_date, goal_words, id, is_finished, start_date, title, description, written_words, words_per_day, is_expired FROM goals WHERE book_id=$1`
	query += utils.GetAllowedOrderBy(params.OrderBy)

	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)

	goals = &[]models.GoalsModel{}

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
