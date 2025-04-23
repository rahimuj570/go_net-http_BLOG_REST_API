package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Port string
var Db_host string
var Db_port string
var Db_user string
var DB_Pass string
var Db_name string
var Jwt_secret string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}
	Port = os.Getenv("PORT")
	Db_host = os.Getenv("localhost")
	Db_name = os.Getenv("DB_NAME")
	Db_port = os.Getenv("DB_PORT")
	Db_user = os.Getenv("DB_USER")
	DB_Pass = os.Getenv("DB_PASS")
	Jwt_secret = os.Getenv("JWT_SECRET")
}
