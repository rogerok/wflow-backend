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
		return nil, fmt.Errorf(fmt.Sprintf("%s. Failed conntect to DB", err.Error()))
	}

	return db, nil
}

func CloseConnectionToDb(db *pgx.Conn, ctx context.Context) {
	if db == nil {
		fmt.Println("Database connection is nil, nothing to close")
		return
	}

	err := db.Close(ctx)
	if err != nil {
		fmt.Printf("%s. Failed to close", err.Error())
	}

}
