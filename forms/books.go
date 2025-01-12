package forms

import "github.com/google/uuid"

type BookCreateForm struct {
	Name        string    `json:"name" validate:"required,min=1,max=255"`
	Description *string   `json:"description" validate:"omitempty,min=2,max=255"`
	UserId      uuid.UUID `json:"userId" validate:"required"`
}

func (f *BookCreateForm) Validate() error {
	return validateStruct(f, validate)
}
