package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string
var EVOURL string
var EVOTOKEN string
var WUZURL string

func Load() error {
	err := godotenv.Load()
	if err != nil {
		// Just log warning instead of fatal error
		log.Println("Warning: Error loading .env file, using environment variables")
	}

	DBURL = os.Getenv("DB_URL")
	EVOURL = os.Getenv("EVO_URL")
	EVOTOKEN = os.Getenv("EVO_TOKEN")
	WUZURL = os.Getenv("WUZ_URL")

	// Check for essential environment variables
	if DBURL == "" || EVOURL == "" || EVOTOKEN == "" || WUZURL == "" {
		return fmt.Errorf("essential environment variables missing")
	}

	return nil
}
