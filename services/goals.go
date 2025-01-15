package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/utils"
)

type GoalsService interface {
	CreateGoal(goal *forms.GoalCreateForm) (id *string, err error)
}

type goalsService struct {
	r repositories.GoalsRepository
}

func NewGoalsService(r repositories.GoalsRepository) GoalsService {
	return &goalsService{
		r: r,
	}
}

func (s *goalsService) CreateGoal(goal *forms.GoalCreateForm) (id *string, err error) {

	goalData := models.GoalsModel{
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

	id, err = s.r.CreateGoal(&goalData)

	if err != nil {
		return nil, err
	}

	return id, nil

}
