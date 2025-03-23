package models

import (
	"github.com/google/uuid"
	"time"
)

type GoalStats struct {
	WrittenWords float64 `json:"writtenWords" db:"written_words"`
	WordsPerDay  float64 `json:"wordsPerDay" db:"words_per_day"`
}

type Goals struct {
	BookId       string    `json:"bookId" db:"book_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	EndDate      time.Time `json:"endDate" db:"end_date"`
	GoalWords    float64   `json:"goalWords" db:"goal_words"`
	Id           uuid.UUID `json:"id" db:"id"`
	IsFinished   bool      `json:"isFinished" db:"is_finished"`
	StartDate    time.Time `json:"startDate" db:"start_date"`
	Title        string    `json:"title" db:"title"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	UserId       string    `json:"-" db:"user_id"`
	Description  *string   `json:"description" db:"description"`
	IsExpired    bool      `json:"isExpired" db:"is_expired"`
	WrittenWords float64   `json:"writtenWords" db:"written_words"`
	WordsPerDay  float64   `json:"wordsPerDay" db:"words_per_day"`
}

type GoalsQueryParams struct {
	PaginationQuery
	OrderBy string    `json:"orderBy" default:"createdAt desc"`
	BookId  *string   `json:"bookId"`
	UserId  uuid.UUID `json:"-"`
}

type GoalUpdateResponse struct {
	GoalWords   float64 `json:"goalWords" db:"goal_words"`
	WordsPerDay float64 `json:"wordsPerDay" db:"words_per_day"`
}
