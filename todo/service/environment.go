package service

import "os"

func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); len(val) > 0 {

		return val

	}

	return fallback
}
