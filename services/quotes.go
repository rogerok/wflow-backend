package services

import (
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
)

type QuotesService interface {
	GetRandom() (*models.Quotes, error)
}

type quotesService struct {
	quotesRepository repositories.QuotesRepository
}

func NewQuotesService(quotesRepository repositories.QuotesRepository) QuotesService {
	return &quotesService{quotesRepository: quotesRepository}
}

func (s *quotesService) GetRandom() (*models.Quotes, error) {
	return s.quotesRepository.GetRandom()
}
