package config

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Db struct {
	Host     string `envconfig:"DB_HOST" validate:"required"`
	DbName   string `envconfig:"DB_NAME" validate:"required"`
	User     string `envconfig:"DB_USER" validate:"required"`
	Password string `envconfig:"DB_PASSWORD"`
	Port     uint16 `envconfig:"DB_PORT" validate:"min=1,max=65535"`
}

type App struct {
	HttpPort string `envconfig:"APP_HTTP_PORT" validate:"required"`
}

type Config struct {
	App App
	Db  Db
}

func Parse() (*Config, error) {
	var cnf Config

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, errors.Join(err, errors.New("read .env file"))
	}

	if err := envconfig.Process("", &cnf); err != nil {
		return nil, errors.Join(err, errors.New("read environment"))
	}

	if err := validator.New().Struct(&cnf); err != nil {
		return nil, errors.Join(err, errors.New("validate config"))
	}

	return &cnf, nil
}
