package env

var env_map = map[string]string{}

// returns env variable for key
func GetEnv(key string) (string, bool) {
	val, ok := env_map[key]
	return val, ok
}

// sets key to value in env map
func SetEnv(key string, value string) {
	env_map[key] = value
}
