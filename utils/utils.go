package utils

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func HandlePagination(page int, perPage int) (offset int, selectAll bool) {
	if page >= 1 && perPage > 0 {
		offset = perPage * (page - 1)
		selectAll = false
	} else {
		offset = 0
		selectAll = true
	}

	return offset, selectAll
}

func PasswordValidator(fl validator.FieldLevel) (check bool) {

	patterns := []string{
		`[0-9]`,                          // At least one digit
		`[a-z]`,                          // At least one lowercase letter
		`[A-Z]`,                          // At least one uppercase letter
		`[!@#$%^&*()\-+}{'"[:;>.?/_~\|]`, // At least one special character
	}

	password := fl.Field().String()

	for _, pattern := range patterns {
		match, _ := regexp.MatchString(pattern, password)
		if !match {
			return false
		}
	}

	return true

}
