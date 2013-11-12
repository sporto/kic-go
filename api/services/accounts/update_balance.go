package accounts

import (
	"github.com/sporto/kic/api/models"
	r "github.com/dancannon/gorethink"
)

type UpdateBalanceServ struct {
}

// Updates the balance
// And saves the account
func (serv *UpdateBalanceServ) Run(dbSession *r.Session, account *models.Account) (err error) {

	interest, err := new(CalculateInterestToPayServ).Run(*account)
	if err != nil {
		return
	}

	account.CurrentBalance += interest

	updateServ := new(UpdateServ)
	err = updateServ.Run(dbSession, *account)
	if err != nil {
		return
	}

	return
}
