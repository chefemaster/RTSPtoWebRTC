package main

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
}

func NewEnv() *Env {
	return &Env{}
}

func (element *Env) Get(key string) string {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}
