package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
)

type Reports interface {
	Create(report *forms.ReportCreateForm) (id *string, err error)
	GetListByGoalId(params *models.ReportsQueryParams) (reports *[]models.ReportsModel, err error)
}

type reportsService struct {
	rReports repositories.ReportsRepository
	rGoals   repositories.GoalsRepository
}

func NewReportsService(rReports repositories.ReportsRepository, rGoals repositories.GoalsRepository) Reports {
	return &reportsService{rReports: rReports, rGoals: rGoals}
}

func mapFormToReportModel(report *forms.ReportCreateForm) *models.ReportsModel {
	return &models.ReportsModel{
		BookId:      report.BookId,
		GoalId:      report.GoalId,
		WordsAmount: report.WordsAmount,
		Title:       report.Title,
		Description: report.Description,
		UserId:      report.UserId,
	}
}

func (s *reportsService) Create(report *forms.ReportCreateForm) (id *string, err error) {

	reportData := mapFormToReportModel(report)

	id, err = s.rReports.Create(reportData)

	if err != nil {
		return nil, err
	}

	err = s.rGoals.RecalculateGoal(reportData.WordsAmount, reportData.GoalId)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *reportsService) GetListByGoalId(params *models.ReportsQueryParams) (reports *[]models.ReportsModel, err error) {
	reports, err = s.rReports.GetListByGoalId(params)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
