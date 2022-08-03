package configs

import (
	"go-clean-architecture/src/application/shared"
	"os"
)

func Get(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found && !shared.IsNullOrEmpty(value) {
		return value
	}
	return defaultValue
}
