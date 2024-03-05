package config

import (
	"os"

	"github.com/joho/godotenv"
)

type IEnv interface {
	Load() error // Load env file
}

// Env empty structen
type Env struct{}

// Create new env obj
func NewEnvConfig() IEnv {
	return &Env{}
}

func (e *Env) Load() error {
	_, dockerEnv := os.LookupEnv("DOCKER_ENV")

	var envFile string
	if dockerEnv {
		envFile = "config/.env.docker"
	} else {
		envFile = "config/.env"
	}

	return godotenv.Load(envFile)
}
