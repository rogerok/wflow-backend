package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load environment %s", err.Error())
		return
	}
}
