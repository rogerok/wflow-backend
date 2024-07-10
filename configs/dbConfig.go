package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type PostgresConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
	SslMode  string
}

func GetDbConfig() PostgresConfig {
	return PostgresConfig{
		Driver:   "postgres",
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

func ConnectToDb() (db *sqlx.DB, err error) {
	config := GetDbConfig()
	db, err = sqlx.Connect(config.Driver, config.String())

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("%s. Failed conntect to DB", err.Error()))
	}

	return db, nil
}

func CloseConnectionToDb(db *sqlx.DB) {
	if db == nil {
		fmt.Println("Database connection is nil, nothing to close")
		return
	}

	err := db.Close()
	if err != nil {
		fmt.Printf("%s. Failed to close", err.Error())
	}

}
