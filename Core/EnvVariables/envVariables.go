package Security

import (
	"os"
)

func GoDotEnvVariable(key string) string {

	// load .env file

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	return os.Getenv(key)
}
