package main

import (
	"encoding/json"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/gorilla/mux"
	"github.com/sporto/kic-api/models"
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

func main() {

	initDb()

	r := mux.NewRouter()

	r.HandleFunc("/accounts", accountsIndex)
	r.HandleFunc("/accounts/{id}", accountsShow)

	http.Handle("/", r)

	// http.HandleFunc("/new", insertBookmark)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}

// func insertBookmark(res http.ResponseWriter, req *http.Request) {
// 	session := sessionArray[0]

// 	b := new(Bookmark)
// 	json.NewDecoder(req.Body).Decode(b)

// 	var response r.WriteResponse

// 	err := r.Table("bookmarks").Insert(b).Run(session).One(&response)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	data, _ := json.Marshal("{'bookmark':'saved'}")
// 	res.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	res.Write(data)
// }

func accountsIndex(res http.ResponseWriter, req *http.Request) {
	session := sessionArray[0]
	var response []models.Account

	err := r.Table("accounts").Run(session).All(&response)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}

func accountsShow(res http.ResponseWriter, req *http.Request) {
	session := sessionArray[0]
	var response models.Account

	vars := mux.Vars(req)
	id := vars["id"]

	err := r.Table("accounts").Get(id).Run(session).One(&response)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}
