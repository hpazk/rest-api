package config

import "os"

type DbConnection struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
}

func DbConfig() DbConnection {
	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return dbConfig
}
