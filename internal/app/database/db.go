package database

var db = map[string]string{}

func Set(key, value string) {
	db[key] = value
}

func Get(key string) (string, bool) {
	val, exists := db[key]
	return val, exists
}
