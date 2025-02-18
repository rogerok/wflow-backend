package utils

import (
	"fmt"
	"github.com/rogerok/wflow-backend/constants"
	"golang.org/x/crypto/bcrypt"
	"math"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hash, password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
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

func GetAllowedOrderBy(order string) string {
	if constants.AllowedOrderBy[order] == "" {
		return " ORDER BY created_at desc "
	}

	return " ODER BY " + constants.AllowedOrderBy[order]
}

func GetOffsetLimitQuery(perPage int, offset int) string {
	return fmt.Sprintf(" LIMIT %v OFFSET %v", perPage, offset)
}

func CalculateWordsPerDay(totalWords int, totalDays int) float64 {
	wordsPerDay := float64(totalWords) / float64(totalDays)
	roundedWordsPerDay := math.Round(wordsPerDay*10) / 10

	return roundedWordsPerDay

}
