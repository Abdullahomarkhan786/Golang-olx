package config

import (
	"os"

	"github.com/joho/godotenv"
)

// to group these variables together we create a struct
type Config struct {
	Port string
	Env  string
}

func MustLoad() Config { //return the struct
	godotenv.Load()
	//check if the variable of env exist or not
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	env := os.Getenv("Env")
	if env == "" {
		panic("Env is required")
	}
	return Config{
		Port: port,
		Env:  env,
	}
}
