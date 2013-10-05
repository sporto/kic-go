package main

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/models"
	"log"
	"time"
)

func main() {
	writeOne()
}

func writeOne() {
	session, err := r.Connect("localhost:28015", "kic")
	if err != nil {
		fmt.Println("error connecting:", err)
		return
	}

	t := time.Now()
	account := models.Account{Name: "Nico", LastInterestPaid: t}

	var response r.WriteResponse
	err = r.Table("accounts").Insert(account).Run(session).One(&response)

	if err != nil {
		log.Fatal(err)
		return
	}

}