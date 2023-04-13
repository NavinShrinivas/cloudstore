package loader

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
)

func CheckAndSetVariables() {
	envRequired := []string{"PORT", "CORS_ORIGIN", "JWT_SECRET", "JWT_LIFETIME", "DATABASE_USERNAME", "DATABASE_PASSWORD", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME"}

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

	lifeTime, err := strconv.ParseInt(os.Getenv("JWT_LIFETIME"), 10, 64)
	if err != nil {
		log.Panic("Error parsing JWT_LIFETIME")
	}
	if lifeTime <= 0 {
		log.Panic("Invalid JWT_LIFETIME")
	}
}
