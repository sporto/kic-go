package accounts

import (
	// "fmt"
	"github.com/sporto/kic/api/models"
	r "github.com/dancannon/gorethink"
)

type UpdateBalanceServ struct {
}

func (serv *UpdateBalanceServ) Run(dbSession *r.Session, account models.Account) (err error) {
	// dur, err := new(CalculateInterestDurationToPayServ).Run(account)
	// if err != null {
	// 	return
	// }

	// interest := misc.CalculateInterest(account.CurrentBalance, days, 3.5)
	// fmt.Println("Interest paid %v", int)
	return
}
