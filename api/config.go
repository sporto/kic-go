package api

import (
	// "github.com/joho/godotenv"
	"os"
)


func GetConfigVar(name string) string {
	env := os.Getenv("ENV")

	if os.Getenv("WERCKER") == "true" {
		env = "wercker"
	}

	switch name {
		case "DB_HOST":
			return getDbHost(env)
		case "DB_NAME":
			return getDbName(env)
		default:
			return os.Getenv("KIC_DEV_DB_HOST")
	}
}

func getDbHost(env string) string {
	switch env {
	case "wercker":
		return os.Getenv("WERCKER_RETHINKDB_URL")
	case "test":
		return os.Getenv("KIC_TEST_DB_HOST")
	case "prod":
		return os.Getenv("KIC_PROD_DB_HOST")
	default:
		return os.Getenv("KIC_DEV_DB_HOST")
	}
}

func getDbName(env string) string {
	switch env {
	case "wercker":
		return "kic_test"
	case "test":
		return os.Getenv("KIC_TEST_DB_NAME")
	case "prod":
		return os.Getenv("KIC_PROD_DB_NAME")
	default:
		return os.Getenv("KIC_DEV_DB_NAME")
	}
}