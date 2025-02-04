package models

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Description *string   `json:"description" db:"description"`
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"book_name"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	UserId      string    `json:"-" db:"user_id"`
}

type BooksQueryParams struct {
	PaginationQuery
	OrderBy string    `json:"orderBy" default:"createdAt desc"`
	UserId  uuid.UUID `json:"-"`
}
