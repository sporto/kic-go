package api

import (
	r "github.com/dancannon/gorethink"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getDbSession() (dbSession *r.Session, err error) {

	address, database := getDbConf()

	// global setup
	dbSession, err = r.Connect(map[string]interface{}{
		"address":  address,
		"database": database,
	})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return
}

func initDb(dbSession *r.Session) error {

	_, database := getDbConf()

	_, err := r.DbCreate(database).Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db(database).TableCreate("accounts").Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db(database).TableCreate("transactions").Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	return nil
}


func getDbConf() (address string, database string) {
	if os.Getenv("WERCKER") == "true" {
		address = os.Getenv("WERCKER_RETHINKDB_URL")
		database = "kic_test"
	} else {
		address = os.Getenv("TEST_DB_HOST")
		database = os.Getenv("TEST_DB_NAME")
	}
	return
}

func StartDb(pathToRoot string) (dbSession *r.Session, err error) {
	// load env
	err = godotenv.Load(pathToRoot + ".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbSession, err = getDbSession()
	if err != nil {
		log.Println(err)
	}

	err = initDb(dbSession)
	if err != nil {
		log.Println(err)
	}

	return
}
