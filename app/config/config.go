package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
	DBNAME string
}

type Config struct {
	DBConfig
}

func (cfg *Config) readConfig() {

	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal(err)
	}

	cfg.DBConfig = DBConfig{
		DBUSER: os.Getenv("DBUSER"),
		DBPASS: os.Getenv("DBPASS"),
		DBHOST: os.Getenv("DBHOST"),
		DBPORT: os.Getenv("DBPORT"),
		DBNAME: os.Getenv("DBNAME"),
	}

}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.readConfig()

	return cfg
}
