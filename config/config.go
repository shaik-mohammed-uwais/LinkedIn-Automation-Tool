package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration required by the bot.
// Centralizing config makes the app easier to understand and maintain.
type Config struct {
	Email    string
	Password string

	Headless bool
	LogLevel string

	DailyConnectLimit int
}

// LoadConfig loads environment variables (including .env),
// applies defaults, and validates required fields.
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (no error if missing)
	_ = godotenv.Load()

	cfg := &Config{
		Email:    os.Getenv("LINKEDIN_EMAIL"),
		Password: os.Getenv("LINKEDIN_PASSWORD"),

		LogLevel: getString("LOG_LEVEL", "info"),
		Headless: getBool("HEADLESS", false),

		DailyConnectLimit: getInt("DAILY_CONNECTION_LIMIT", 20),
	}

	// Validate required credentials
	if cfg.Email == "" || cfg.Password == "" {
		return nil, errors.New(
			"missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD in environment",
		)
	}

	return cfg, nil
}

// getString returns the value of an environment variable
// or a fallback value if the variable is not set.
func getString(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// getBool safely parses a boolean environment variable.
// Falls back to the provided default on error.
func getBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}

// getInt safely parses an integer environment variable.
// Falls back to the provided default on error.
func getInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
