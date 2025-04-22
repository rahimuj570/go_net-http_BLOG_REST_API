package config

import "github.com/joho/godotenv"

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}
	return nil
}
