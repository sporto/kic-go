package accounts

import (
	"fmt"
	"github.com/sporto/kic/api/models"
)

func UpdateAccount(account models.Account) {
	// check when was the last time interest was paid
	// loop through each day
	// - pay interest
	// - update the balance
	fmt.Println(account)

	PayInterest(account)
}
