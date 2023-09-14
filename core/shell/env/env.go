package env

// contains all env variables
var env_map = map[string]string{
	"PWD":                    "/",
	"GPNZR_PROMPT_SHORT_PWD": "",
}

// returns value for key of 'env_map', false if not found
func GetEnv(key string) (string, bool) {
	val, ok := env_map[key]
	return val, ok
}

// returns true value for key of 'env_map' equal to 1, false everything else and not found
func GetEnvBool(key string) (r bool) {
	r = false
	val, ok := env_map[key]

	if !ok {
		return
	}

	if val == "1" {
		return true
	}

	return
}

// sets key to value in 'env_map'
func SetEnv(key string, value string) {
	env_map[key] = value
}
