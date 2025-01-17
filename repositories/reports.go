package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type ReportsRepository interface {
	Create(report *models.ReportsModel) (id *string, err error)
	GetListByGoalId(params *models.ReportsQueryParams) (reports *[]models.ReportsModel, err error)
}

type reportsRepository struct {
	db *sqlx.DB
}

func NewReportsRepository(db *sqlx.DB) ReportsRepository {
	return &reportsRepository{db: db}
}

func (r *reportsRepository) Create(report *models.ReportsModel) (id *string, err error) {
	query := `INSERT INTO reports (book_id, goal_id, user_id, words_amount, title, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = r.db.QueryRow(query, report.BookId, report.GoalId, report.UserId, report.WordsAmount, report.Title, report.Description).Scan(&id)

	if err != nil {
		return nil, err
	}

	return id, nil

}

func (r *reportsRepository) GetListByGoalId(params *models.ReportsQueryParams) (reports *[]models.ReportsModel, err error) {
	query := `SELECT words_amount, title, description FROM reports WHERE goal_id=$1`

	query += utils.GetAllowedOrderBy(params.OrderBy)

	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)

	reports = &[]models.ReportsModel{}

	if selectAll {
		err = r.db.Select(reports, query, params.GoalId)
	} else {
		query += utils.GetOffsetLimitQuery(params.PerPage, offset)
		err = r.db.Select(reports, query, params.GoalId)
	}

	if err != nil {
		return nil, err
	}

	return reports, nil

}
