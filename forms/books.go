package forms

type BookCreateForm struct {
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description *string `json:"description" validate:"omitempty,min=2,max=255"`
	UserId      string  `json:"-"`
}

func (f *BookCreateForm) Validate() error {
	return validateStruct(f, validate)
}
