package tasks

import (
	// r "github.com/dancannon/gorethink"
	"github.com/chuckpreslar/gofer"
	"github.com/sporto/kic/api"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
	"log"
)

var TaskOne = gofer.Register(gofer.Task{
	Namespace:   "accounts",
	Label:       "create",
	Description: "Create Nico's account",
	Action: func(arguments ...string) (err error) {
		log.Println("accounts:create")

		dbSession, err := api.StartDb("./")
		if err != nil {
			log.Fatal(err)
		}

		account := &models.Account{
			Name: "Nico",
			LastInterestPaid: time.Now(),
		}

		serv := &accounts.CreateServ{}
		_, err = serv.Run(dbSession, *account)
		if err != nil {
			log.Fatal(err)
		}

		return
	},
})