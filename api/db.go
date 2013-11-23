package api

import (
	r "github.com/dancannon/gorethink"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDbSession(path string) (dbSession *r.Session, err error) {

	// load env
	err = godotenv.Load(path + ".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbAddress := ""
	dbName := ""

	// log.Println("DB Host", os.Getenv("TEST_DB_HOST"))
	// log.Println("DB Name", os.Getenv("TEST_DB_NAME"))

	if os.Getenv("WERCKER") == "true" {
		dbAddress = os.Getenv("WERCKER_RETHINKDB_URL")
		dbName = os.Getenv("kic_test")
	} else {
		dbAddress = os.Getenv("TEST_DB_HOST")
		dbName = os.Getenv("TEST_DB_NAME")
	}

	// global setup
	dbSession, err = r.Connect(map[string]interface{}{
		"address":  dbAddress,
		"database": dbName,
	})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return
}
