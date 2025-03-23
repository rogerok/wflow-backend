package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/utils"
	"time"
)

type GoalsService interface {
	Create(goal *forms.GoalCreateForm) (id *string, err error)
	Edit(goal *forms.GoalEditForm) (goalStats *models.GoalUpdateResponse, err error)
	Delete(goalId string, userId string) (status bool, err error)
	GetList(params *models.GoalsQueryParams) (goals *[]models.Goals, err error)
	GetById(id string) (goal *models.Goals, err error)
	RecalculateGoals()
}

type goalsService struct {
	r repositories.GoalsRepository
}

func NewGoalsService(r repositories.GoalsRepository) GoalsService {
	return &goalsService{
		r: r,
	}
}

func mapFormToModel(goal *forms.GoalCreateForm) *models.Goals {
	start := time.Date(goal.StartDate.Year(), goal.StartDate.Month(), goal.StartDate.Day(), 0, 0, 0, 0, goal.StartDate.Location())
	end := time.Date(goal.EndDate.Year(), goal.EndDate.Month(), goal.EndDate.Day(), 0, 0, 0, 0, goal.EndDate.Location())

	days := int(end.Sub(start).Hours()/24) + 1

	return &models.Goals{
		BookId:       goal.BookId,
		EndDate:      goal.EndDate,
		GoalWords:    goal.GoalWords,
		IsFinished:   false,
		StartDate:    goal.StartDate,
		Title:        goal.Title,
		UserId:       goal.UserId,
		Description:  goal.Description,
		WrittenWords: 0,
		WordsPerDay:  utils.CalculateWordsPerDay(goal.GoalWords, days),
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

func (s *goalsService) Edit(goal *forms.GoalEditForm) (goalStats *models.GoalUpdateResponse, err error) {

	return s.r.Edit(goal)

}

func (s *goalsService) Delete(goalId string, userId string) (status bool, err error) {

	return s.r.Delete(goalId, userId)
}

func (s *goalsService) GetList(params *models.GoalsQueryParams) (goals *[]models.Goals, err error) {
	goals, err = s.r.GetList(params)

	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (s *goalsService) GetById(id string) (goal *models.Goals, err error) {
	goal, err = s.r.GetById(id)

	if err != nil {
		return nil, err
	}

	return goal, nil
}

func (s *goalsService) RecalculateGoals() {
	s.r.RecalculateGoals()
}
