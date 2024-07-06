package configs

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
	SslMode  string
}

func GetDbConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("POSTGRES_DB_NAME"),
		SslMode:  os.Getenv("SSL_MODE"),
	}
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode)
}

func ConnectToDb() (db *pgx.Conn, err error) {
	confString := GetDbConfig().String()
	db, err = pgx.Connect(context.Background(), confString)

	if err != nil {
		panic(fmt.Sprintf("%s. Failed conntect to DB", err.Error()))
	}

	return db, nil
}
