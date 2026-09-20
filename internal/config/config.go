package config

import (
	"os"

	"github.com/joho/godotenv"
)

// to store config and group multiple properties use struct
type Config struct {
	Port string
	Env  string
}

// must prefix must be loaded it loads the config variables
// config must be there and it returns config
func MustLoad() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is required")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("env is required")
	}
	return Config{
		Port: port,
		Env:  env,
	}

}
