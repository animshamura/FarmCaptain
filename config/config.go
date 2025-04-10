package config

import "os"

// GetEnv retrieves the value of an environment variable or returns a default value.
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
