package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type GoalsRepository interface {
	CreateGoal(goal *models.GoalsModel) (id *string, err error)
	GetGoalById(id string) (goal *models.GoalsModel, err error)
}

type goalsRepository struct {
	db *sqlx.DB
}

func NewGoalsRepository(db *sqlx.DB) GoalsRepository { return &goalsRepository{db: db} }

func (r *goalsRepository) CreateGoal(goal *models.GoalsModel) (id *string, err error) {

	query := `INSERT INTO goals 
    						(book_id, end_date, goal_words, start_date, title, user_id, description, words_per_day, written_words)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = r.db.QueryRow(query, goal.BookId, goal.EndDate, goal.GoalWords, goal.StartDate, goal.Title, goal.UserId, goal.Description, goal.WordsPerDay, goal.WrittenWords).Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *goalsRepository) GetGoalById(id string) (goal *models.GoalsModel, err error) {
	query := `SELECT * FROM goals WHERE  id=$1`

	err = r.db.Get(query, id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}
