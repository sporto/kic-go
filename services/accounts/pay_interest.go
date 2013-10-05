package accounts

import (
	"fmt"
	"github.com/sporto/kic/models"
	"github.com/sporto/kic/services/misc"
)

func PayInterest(account models.Account) {
	fmt.Println("pay interest")

	days := 1.0

	// dur := CalculateInterestDurationToPay(account)

	int := misc.CalculateInterest(account.CurrentBalance, days, 3.5)
	fmt.Println("Interest paid %v", int)

}
