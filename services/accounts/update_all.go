package accounts

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic-api/models"
)

func UpdateAll() {
	session, err := r.Connect("localhost:28015", "kic")
	if err != nil {
		fmt.Println("error connecting:", err)
		return
	}

	// If we want to iterate over each result individually, we can use the rows
	// object as an iterator
	rows := r.Table("accounts").Run(session)
	for rows.Next() {
		var account models.Account
		if err = rows.Scan(&account); err != nil {
			fmt.Println("err:", err)
			break
		}
		UpdateAccount(account)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("err:", err)
	}
}
