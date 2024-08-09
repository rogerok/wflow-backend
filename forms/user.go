package forms

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

type Pseudonym struct {
	FirstName *string `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"lastName" validate:"omitempty,min=2,max=50"`
}

type Social struct {
	Instagram *string `json:"instagram" validate:"omitempty,url"`
	Telegram  *string `json:"telegram" validate:"omitempty,url"`
	TikTok    *string `json:"tiktok" validate:"omitempty,url"`
	Vk        *string `json:"vk" validate:"omitempty,url"`
}

type User struct {
	Email       string     `json:"email" validate:"required,email,max=255"`
	FirstName   string     `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastName    *string    `json:"lastName" validate:"omitempty,min=2,max=50"`
	MiddleName  *string    `json:"middleName" validate:"omitempty,min=2,max=50"`
	Password    string     `json:"-" validate:"required,min=8,max=255"`
	Pseudonym   Pseudonym  `json:"pseudonym" validate:"dive"`
	SocialLinks Social     `json:"socialLinks" validate:"dive"`
	BornDate    *time.Time `json:"bornDate" db:"omitempty,datetime=2006-01-02"`
}

func PasswordValidator(fl validator.FieldLevel) (check bool) {
	checks := []string{
		"[0-9]", // Checks for at least one digit
		"[a-z]", // Checks for at least one lowercase letter
		"[A-Z]", // Checks for at least one uppercase letter
		"[!@#$%^&*()\\-\\+}{'\";:.,></\\?\\|_=`~]", // Checks for at least one special character
	}

	checks = append(checks, fmt.Sprintf(`^(%s)+$`), strings.Join(checks, "|"))

	for _, c := range checks {
		if check = regexp.MustCompile(c).MatchString(fl.Field().String()); !check {
			return
		}
	}

	return true

}
