package main

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/api/controllers"
	"github.com/stretchr/goweb"
	"log"
	"net/http"
)

var sessionArray []*r.Session

func initDb() {
	// session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
	session, err := r.Connect("localhost:28015", "kic")
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer session.Close()

	err = r.DbCreate("kic").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	err = r.TableCreate("accounts").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	sessionArray = append(sessionArray, session)
}

func mapRoutes() {
	accountsController := &controllers.Accounts{DbSession: sessionArray[0]}

	goweb.MapController("/api/accounts", accountsController)

	goweb.MapStatic("/public", "src")
	goweb.MapStaticFile("/index.html", "src/index.html")
}

func main() {
	initDb()
	mapRoutes()

	http.Handle("/", goweb.DefaultHttpHandler())

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}
