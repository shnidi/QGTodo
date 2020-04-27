package goEnv

import (
	"fmt"
	"os"
	"strconv"
)

func StrictGetEnv(key string) (v string, err error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("Empty env")
}
func StrictGetEnvToI(key string) (v int, err error) {
	if v := os.Getenv(key); v != "" {
		return strconv.Atoi(v)
	}
	return 0, fmt.Errorf("Empty env")
}
