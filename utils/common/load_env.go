package common

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error Load env file: %s", err.Error())
	}
	return nil
}
