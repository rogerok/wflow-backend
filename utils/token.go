package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/responses"
	"os"
	"strings"
	"time"
)

func CreateToken(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{"sub": id, "iss": "wflow", "exp": time.Now().Add(time.Minute * 60).Unix(), "iat": time.Now().Unix()})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return token, err

}

func CreateRefreshTokenExpTime() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}

func CreateRefreshToken(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": id, "exp": CreateRefreshTokenExpTime().Unix()})

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

func ExtractTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	token := ctx.Get("Authorization")

	if !strings.HasPrefix(token, "Bearer ") {
		return "", GetInvalidTokenError(ctx)
	}

	return strings.TrimPrefix(token, "Bearer "), nil
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
	token, err := ExtractTokenFromHeader(ctx)

	if err != nil {
		return "", GetInvalidTokenError(ctx)
	}

	return GetSubjectFromToken(token)
}

func GetSubjectUuidFromHeaderToken(ctx *fiber.Ctx) (uuid.UUID, error) {
	subject, err := GetSubjectFromHeaderToken(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	parsed, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, GetInvalidTokenError(ctx)
	}

	return parsed, nil
}
