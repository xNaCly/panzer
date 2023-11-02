// provides abstractions for environment interactions, such as getting and
// setting environment variables - even supporting boolean environment
// variables
package env

import "os"

// returns true if environment variable 'key' is equal to 1, false for all
// other values and not environment variable empty / not found
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
