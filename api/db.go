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
		log.Fatal("Error loading .env file")
	}

	log.Println("DB Host", os.Getenv("TEST_DB_HOST"))
	log.Println("DB Name", os.Getenv("TEST_DB_NAME"))

	// global setup
	dbSession, err = r.Connect(map[string]interface{}{
		"address":  os.Getenv("TEST_DB_HOST"),
		"database": os.Getenv("TEST_DB_NAME"),
	})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return
}
