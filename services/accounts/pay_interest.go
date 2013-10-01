package accounts

import (
	"fmt"
	"github.com/sporto/kic/models"
	"github.com/sporto/kic/services/misc"
)

func PayInterest(account models.Account) {
	fmt.Println("pay interest")

	days := 1.0

	if account.LastInterestPaid == "" {
		// calc the days
	} else {
		fmt.Println("Last interest paid on %v", account.LastInterestPaid)
		// pay interest from yesterday
	}

	int := misc.CalculateInterest(account.Balance, days, 3.5)
	fmt.Println("Interest paid %v", int)

}
