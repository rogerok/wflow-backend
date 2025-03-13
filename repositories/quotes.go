package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type QuotesRepository interface {
	GetRandom() (*models.Quotes, error)
}

type quotesRepository struct {
	db *sqlx.DB
}

func NewQuotesRepository(db *sqlx.DB) QuotesRepository {
	return &quotesRepository{db: db}
}

func (r *quotesRepository) GetRandom() (*models.Quotes, error) {
	quote := &models.Quotes{}

	query := `SELECT * FROM writing_quotes ORDER BY RANDOM() LIMIT 1`

	err := r.db.Get(quote, query)

	if err != nil {
		return nil, err
	}

	return quote, nil
}
