package config

import "github.com/joho/godotenv"

type Env struct{}

func (e *Env) Load() error {
	return godotenv.Load("config/.env")
}
