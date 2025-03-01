package forms

import "github.com/rogerok/wflow-backend/forms/validators"

type BookCreateForm struct {
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description *string `json:"description" validate:"omitempty,min=2,max=255"`
	UserId      string  `json:"-"`
}

func (f *BookCreateForm) Validate() error {
	return validators.ValidateStruct(f, validators.Validate)
}
