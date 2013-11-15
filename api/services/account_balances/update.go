package account_balances

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/transactions"
	"github.com/sporto/kic/api/services/accounts"
	"time"
)

type UpdateServ struct {
}

// Updates the balance
// And saves the account
func (serv *UpdateServ) Run(dbSession *r.Session, accountIn models.Account) (accountOut models.Account, transactionOut models.Transaction, err error) {

	interest, err := new(accounts.CalculateInterestToPayServ).Run(accountIn)
	if err != nil {
		return
	}

	// create a transaction
	transaction := &models.Transaction{
		AccountId: accountIn.Id,
		Credit:    interest,
		Kind:      "interest",
	}
	transaction.Credit = interest
	transaction.Kind = "interest"

	createTransactionServ := &transactions.CreateServ{}
	transactionOut, err = createTransactionServ.Run(dbSession, *transaction)
	if err != nil {
		return
	}

	accountIn.CurrentBalance   += interest
	accountIn.LastInterestPaid = time.Now()

	updateAccountServ := &accounts.UpdateServ{}
	accountOut, err = updateAccountServ.Run(dbSession, accountIn)
	if err != nil {
		return
	}

	return
}
