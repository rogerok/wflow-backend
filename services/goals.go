package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/utils"
)

type GoalsService interface {
	Create(goal *forms.GoalCreateForm) (id *string, err error)
	GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.GoalsModel, err error)
	GetById(id string) (goal *models.GoalsModel, err error)
}

type goalsService struct {
	r repositories.GoalsRepository
}

func NewGoalsService(r repositories.GoalsRepository) GoalsService {
	return &goalsService{
		r: r,
	}
}

func mapFormToModel(goal *forms.GoalCreateForm) *models.GoalsModel {
	return &models.GoalsModel{
		BookId:       goal.BookId,
		EndDate:      goal.EndDate,
		GoalWords:    goal.GoalWords,
		IsFinished:   false,
		StartDate:    goal.StartDate,
		Title:        goal.Title,
		UserId:       goal.UserId,
		Description:  goal.Description,
		WrittenWords: 0,
		WordsPerDay:  utils.CalculateWordsPerDay(goal.GoalWords, int(goal.EndDate.Sub(goal.StartDate).Hours())),
	}
}

func (s *goalsService) Create(goal *forms.GoalCreateForm) (id *string, err error) {
	goalData := mapFormToModel(goal)

	id, err = s.r.Create(goalData)

	if err != nil {
		return nil, err
	}

	return id, nil

}

func (s *goalsService) GetListByBookId(params *models.GoalsQueryParams) (goals *[]models.GoalsModel, err error) {
	goals, err = s.r.GetListByBookId(params)

	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (s *goalsService) GetById(id string) (goal *models.GoalsModel, err error) {
	goal, err = s.r.GetById(id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}