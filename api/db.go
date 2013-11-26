package api

import (
	r "github.com/dancannon/gorethink"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getDbSession() (dbSession *r.Session, err error) {

	address, database := getDbConf()

	if address == "" || database == "" {
		log.Fatal("Invalid db config")
	}

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
	log.Println("initDb")

	_, database := getDbConf()

	log.Println("Database name ", database)

	_, err := r.DbCreate(database).Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	// check db
	// _, err = r.Db(database).Run(dbSession)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	env := os.Getenv("ENV")
	if os.Getenv("WERCKER") == "true" {
		env = "wercker"
	}

	// default to dev
	address = os.Getenv("DEV_DB_HOST")
	database = os.Getenv("DEV_DB_NAME")

	switch env {
		case "wercker":
			address = os.Getenv("WERCKER_RETHINKDB_URL")
			database = "kic_test"
		case "test":
			address = os.Getenv("TEST_DB_HOST")
			database = os.Getenv("TEST_DB_NAME")
		case "prod":
			address = os.Getenv("PROD_DB_HOST")
			database = os.Getenv("PROD_DB_NAME")
	}
	return
}

func StartDb(pathToRoot string) (dbSession *r.Session, err error) {
	log.Println("StartDb")
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
