package config

import (
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

// Decorators
func (e *Env) Load() error {
	return godotenv.Load("config/.env")
}
