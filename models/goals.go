package models

import (
	"github.com/google/uuid"
	"time"
)

type Goals struct {
	BookId       string    `json:"book_id" db:"book_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	EndDate      time.Time `json:"endDate" db:"end_date"`
	GoalWords    int       `json:"goalWords" db:"goal_words"`
	Id           uuid.UUID `json:"id" db:"id"`
	IsFinished   bool      `json:"isFinished" db:"is_finished"`
	StartDate    time.Time `json:"startDate" db:"start_date"`
	Title        string    `json:"title" db:"title"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	UserId       string    `json:"-" db:"user_id"`
	Description  *string   `json:"description" db:"description"`
	WrittenWords int       `json:"writtenWords" db:"written_words"`
	WordsPerDay  float64   `json:"wordsPerDay" db:"words_per_day"`
	IsExpired    bool      `json:"isExpired" db:"is_expired"`
}

type GoalsQueryParams struct {
	PaginationQuery `json:"-"`
	OrderBy         string    `json:"orderBy" default:"createdAt desc"`
	BookId          *string   `json:"bookId"`
	UserId          uuid.UUID `json:"-"`
}
