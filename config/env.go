package config

import (
	"github.com/joho/godotenv"
)

// this function will load the .env file if the GO_ENV environment variable is not set
func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
