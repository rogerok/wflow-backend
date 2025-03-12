package services

import (
	"github.com/google/uuid"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
)

type StatisticsService interface {
	GetUserStatistics(userId uuid.UUID) (*models.UserStatistics, error)
	GetGoalStatistics(goalId string) (*models.GoalStatistics, error)
}

type statisticsService struct {
	r repositories.StatisticsRepository
}

func NewStatisticService(r repositories.StatisticsRepository) StatisticsService {
	return &statisticsService{r: r}
}

func (s *statisticsService) GetUserStatistics(userId uuid.UUID) (*models.UserStatistics, error) {
	return s.r.GetUserStatistics(userId)
}

func (s *statisticsService) GetGoalStatistics(goalId string) (*models.GoalStatistics, error) {
	return s.r.GetGoalStatistics(goalId)
}
