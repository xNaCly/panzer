package env

import "os"

func GetEnv(key string) (string, bool) {
	val := os.Getenv(key)
	return val, val != ""
}

// returns true value for key of 'env_map' equal to 1, false everything else and not found
func GetEnvBool(key string) (r bool) {
	r = false
	val := os.Getenv(key)

	if val == "" {
		return
	}

	if val == "1" {
		return true
	}

	return
}

// sets key to value in 'env_map'
func SetEnv(key string, value string) {
	os.Setenv(key, value)
}
