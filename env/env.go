package env

import (
	"os"
	"strconv"
	"strings"
)

// Get is
func Get(key string, def ...string) string {
	str := strings.TrimSpace(os.Getenv(key))
	if str == "" && len(def) > 0 {
		str = def[0]
	}
	return str
}

// GetInt is
func GetInt(key string, def ...int) int {
	str := strings.TrimSpace(os.Getenv(key))
	if str == "" && len(def) > 0 {
		return def[0]
	}

	in, err := strconv.Atoi(str)
	if err != nil && len(def) > 0 {
		return def[0]
	}

	return in
}
