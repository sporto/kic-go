package tasks

import (
	r "github.com/dancannon/gorethink"
	"github.com/chuckpreslar/gofer"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
	"fmt"
)

var TaskOne = gofer.Register(gofer.Task{
	Namespace:   "accounts",
	Label:       "create",
	Description: "Create Nico's account",
	Action: func(arguments ...string) (err error) {

		account := &models.Account{
			Name: "Nico",
			LastInterestPaid: time.Now(),
		}

		dbSession, _ := r.Connect(map[string]interface{}{
			"address":  "localhost:28015",
			"database": "kic",
		})

		serv := &accounts.CreateServ{}
		_, err = serv.Run(dbSession, *account)
		fmt.Println(err)
		return
	},
})