package forms

type ReportCreateForm struct {
	BookId      string  `json:"bookId" validate:"required"`
	GoalId      string  `json:"goalId" validate:"required"`
	WordsAmount int     `json:"wordsAmount" validate:"required,min=2"`
	Title       string  `json:"title" validate:"required,min=2,max=255"`
	Description *string `json:"description" validate:"omitempty,min=2,max=255"`
	UserId      string  `json:"-"`
}
