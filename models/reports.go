package models

type ReportsModel struct {
	BookId      string  `json:"bookId" db:"book_id"`
	CreatedAt   string  `json:"createdAt" db:"created_at"`
	GoalId      string  `json:"goalId" db:"goal_id"`
	Id          string  `json:"id" db:"id"`
	UpdatedAt   string  `json:"updatedAt" db:"updated_at"`
	WordsAmount int     `json:"wordsAmount" db:"words_amount"`
	Description *string `json:"description" db:"description"`
	Title       string  `json:"title" db:"title"`
	UserId      string  `json:"-" db:"user_id"`
}

type ReportsQueryParams struct {
	PaginationQuery
	OrderBy string `json:"orderBy" default:"createdAt desc"`
	GoalId  string `json:"goalId" validate:"required"`
}
