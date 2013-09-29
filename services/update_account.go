package services

import (
	"fmt"
	"github.com/sporto/kic/models"
)

func UpdateAccount(account models.Account) {
	// check when was the last time interest was paid
	// loop through each day
	// - pay interest
	// - update the balance
	fmt.Println(account)

	if account.LastInterestPaid {
		fmt.Println("INteste paid")
	} else {
		fmt.Println("No")
	}
}
