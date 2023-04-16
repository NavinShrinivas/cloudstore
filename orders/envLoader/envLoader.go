package envLoader

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
)

func CheckAndSetVariables() {
	envRequired := []string{"PORT", "ALLOWED_ORIGINS", "USERHANDLE_ADDRESS", "PRODUCTS_ADDRESS", "DATABASE_USERNAME", "DATABASE_PASSWORD", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME"}

	_, err := os.Stat(".env")
	if err == nil {
		secret, err := godotenv.Read()
		if err != nil {
			log.Panic("Error reading .env file")
		}

		for _, key := range envRequired {
			if secret[key] != "" {
				os.Setenv(key, secret[key])
			}
		}
	}

	for _, key := range envRequired {
		if os.Getenv(key) == "" {
			log.Panic("Environment variable " + key + " not set")
		}
	}
}
