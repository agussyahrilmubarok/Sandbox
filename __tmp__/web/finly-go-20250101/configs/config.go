package configs

import "os"

func GetAppPort() string {
	port := os.Getenv("FINLY_PORT")
	if port == "" {
		return ":8080"
	}

	return port
}

func GetMongoURI() string {
	key := "MONGO_URI"
	if os.Getenv(key) == "" {
		return "mongodb://admin:securepassword123@localhost:27017"
	}

	return os.Getenv(key)
}

func GetSecretKey() string {
	key := "SECRET_KEY"
	if os.Getenv(key) == "" {
		return "secret_key"
	}
	return os.Getenv(key)
}
