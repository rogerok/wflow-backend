package forms

type ReportCreateForm struct {
	BookId      string  `json:"bookId" validate:"required"`
	GoalId      string  `json:"goalId" validate:"required"`
	WordsAmount float64 `json:"wordsAmount" validate:"required,min=2"`
	UserId      string  `json:"-"`
}
