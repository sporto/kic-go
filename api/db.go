package api

import (
	r "github.com/dancannon/gorethink"
	"github.com/joho/godotenv"
	"log"
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
	address = GetConfigVar("DB_HOST")
	database = GetConfigVar("DB_NAME")
	return
}

func StartDb(pathToRoot string) (dbSession *r.Session, err error) {
	log.Println("StartDb")
	// load env
	err = godotenv.Load(pathToRoot + ".env")
	if err != nil {
		log.Println("No .env file")
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
