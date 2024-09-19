package main

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/configs"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/utils"
	"github.com/stretchr/testify/require"

	"log"
	"os"
	"testing"
	"time"
)

type Queries struct {
	db *sqlx.DB
}

var testQueries *Queries

func TestMain(m *testing.M) {
	utils.LoadEnv()

	db, dbError := configs.ConnectToDb()

	testQueries = &Queries{db: db}

	if dbError != nil {
		log.Fatal("cannot connect to db:", dbError)
	}

	err := db.Ping()

	if err != nil {
		return
	}

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	userRepo := repositories.NewUserRepository(testQueries.db)

	user := models.User{
		BornDate:   nil,
		CreatedAt:  time.Time{},
		Email:      "",
		FirstName:  "",
		Id:         uuid.UUID{},
		LastName:   nil,
		MiddleName: nil,
		Password:   "",
		Pseudonym: models.Pseudonym{
			FirstName: nil,
			LastName:  nil,
		},
		SocialLinks: models.Social{
			Instagram: nil,
			Telegram:  nil,
			TikTok:    nil,
			Vk:        nil,
		},
		UpdatedAt: time.Time{},
	}

	id, err := userRepo.CreateUser(&user)

	if err != nil {
		return
	}
	require.NoError(t, err)
	require.NotEmpty(t, id)
}
