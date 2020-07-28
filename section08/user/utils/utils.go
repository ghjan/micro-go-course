package utils

import (
	"os"
)

func GetEnv(key string, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return

}
