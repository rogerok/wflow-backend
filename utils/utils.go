package utils

import (
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
		return "createdAt desc"
	}

	return constants.AllowedOrderBy[order]
}

func CalculateWordsPerDay(totalWords int, totalDays int) float64 {
	wordsPerDay := float64(totalWords) / float64(totalDays)
	roundedWordsPerDay := math.Round(wordsPerDay*10) / 10

	if int(roundedWordsPerDay*float64(totalDays)) < totalWords {
		roundedWordsPerDay += 0.1
	}

	return roundedWordsPerDay

}
