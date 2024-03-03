package config

import (
	"errors"
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	DSN      string
}

func NewDBConfig() (*DBConfig, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return nil, errors.New("POSTGRES_HOST is empty")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return nil, errors.New("POSTGRES_USER is empty")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return nil, errors.New("POSTGRES_PASSWORD is empty")
	}

	dbName := os.Getenv("POSTGRES_DB_NAME")
	if dbName == "" {
		return nil, errors.New("POSTGRES_DB_NAME is empty")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return nil, errors.New("POSTGRES_PORT is empty")
	}

	return &DBConfig{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
		DSN:      fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port),
	}, nil
}
