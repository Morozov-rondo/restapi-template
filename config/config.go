package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
Слой для создания файлов конфигурации.
Загрузка конфигурации из файла .env
*/

var Configs Config

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

func NewConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	Configs.DBUsername = os.Getenv("DB_USERNAME")
	Configs.DBPassword = os.Getenv("DB_PASSWORD")
	Configs.DBHost = os.Getenv("DB_HOST")
	Configs.DBPort = os.Getenv("DB_PORT")
	Configs.DBName = os.Getenv("DB_NAME")

	Configs.ServerPort = os.Getenv("SERVER_PORT")

}
