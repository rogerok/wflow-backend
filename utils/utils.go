package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rogerok/wflow-backend/constants"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/responses"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strings"
	"time"
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

func CreateToken(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{"sub": id, "iss": "wflow", "exp": time.Now().Add(time.Minute * 60).Unix(), "iat": time.Now().Unix()})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return token, err

}

func GetRefreshTokenExpTime() time.Time {
	return time.Now().Add(time.Hour * 24)
}

func CreateRefreshToken(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": id, "exp": GetRefreshTokenExpTime().Unix()})

	rt, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return rt, err
}

func CreateTokenPair(id uuid.UUID) (*responses.TokensModel, error) {

	token, err := CreateToken(id)

	if err != nil {
		return nil, err
	}

	refreshToken, err := CreateRefreshToken(id)

	if err != nil {
		return nil, err
	}

	return &responses.TokensModel{
		Token:        token,
		RefreshToken: refreshToken,
	}, err

}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors_utils.CreateErrorMsg(errors_utils.ErrInvalidToken)
		}

		return []byte(os.Getenv("SECRET_KEY")), nil

	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil

	}

	return nil, err
}

func GetAllowedOrderBy(order string) string {
	if constants.AllowedOrderBy[order] == "" {
		return "createdAt desc"
	}

	return constants.AllowedOrderBy[order]
}

func GetAccessTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	token := ctx.Get("Authorization")

	if token == "" {
		return "", GetInvalidTokenError(ctx)
	}

	parts := strings.Split(token, "Bearer ")

	if len(parts) != 2 {
		return "", GetInvalidTokenError(ctx)
	}

	return parts[1], nil
}

func GetSubjectFromToken(token string) (string, error) {

	parsed, err := ParseToken(token)

	if err != nil {
		return "", err

	}

	subject, err := parsed.GetSubject()

	if err != nil {
		return "", err
	}

	return subject, nil
}

func GetSubjectFromHeaderToken(ctx *fiber.Ctx) (string, error) {
	token, err := GetAccessTokenFromHeader(ctx)

	if err != nil {
		return "", GetInvalidTokenError(ctx)
	}

	subject, err := GetSubjectFromToken(token)

	if err != nil {
		return "", GetInvalidTokenError(ctx)
	}

	return subject, nil
}
