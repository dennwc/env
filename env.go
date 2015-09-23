// Package env provides helpers for getting environment variables with a certain type.
package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Log is a function for logging errors while parsing environment variables.
// Can potentially be changed to panic on errors.
var Log = func(key string, err error) {
	log.Printf("error while parsing %s: %v", key, err)
}

// String gets a string variable from environment. It will use default if variable is empty.
func String(key string, def string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}
	return def
}

// Bool gets a bool variable from environment. It will use default if variable is empty or in wrong format.
func Bool(key string, def bool) bool {
	if s := String(key, ""); s != "" {
		if d, err := strconv.ParseBool(s); err == nil {
			return d
		} else {
			Log(key, err)
		}
	}
	return def
}

// Int gets an int variable from environment. It will use default if variable is empty or in wrong format.
func Int(key string, def int) int {
	if s := String(key, ""); s != "" {
		if d, err := strconv.Atoi(s); err == nil {
			return d
		} else {
			Log(key, err)
		}
	}
	return def
}

// Float64 gets a float64 variable from environment. It will use default if variable is empty or in wrong format.
func Float64(key string, def float64) float64 {
	if s := String(key, ""); s != "" {
		if d, err := strconv.ParseFloat(s, 64); err == nil {
			return d
		} else {
			Log(key, err)
		}
	}
	return def
}

// Duration gets a duration variable from environment. It will use default if variable is empty or in wrong format.
//
// Duration uses time.ParseDuration, so format must follow its rules.
func Duration(key string, def time.Duration) time.Duration {
	if s := String(key, ""); s != "" {
		if d, err := time.ParseDuration(s); err == nil {
			return d
		} else {
			Log(key, err)
		}
	}
	return def
}
