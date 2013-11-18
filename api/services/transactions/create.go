package transactions

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
)

const (
	KindDeposit    = "deposit"
	KindWithdrawal = "withdrawal"
	KindInterest   = "interest"
)

type CreateServ struct {
}

func (serv *CreateServ) Run(dbSession *r.Session, transactionIn models.Transaction) (transactionOut models.Transaction, err error) {

	// fail if transaction is already saved
	if transactionIn.Id != "" {
		err = errors.New("Transaction Id must be nil")
		return
	}

	// fail if not account id provided
	if transactionIn.AccountId == "" {
		err = errors.New("Account Id cannot be nil")
		return
	}

	// fail if no credit or debit provided
	if transactionIn.Credit <= 0 && transactionIn.Debit <= 0 {
		err = errors.New("Credit or Debit must be provided")
		return
	}

	// check that the account exist
	getServ := new(accounts.GetServ)
	account, err := getServ.Run(dbSession, transactionIn.AccountId)
	if err != nil {
		return
	}

	// check that the transaction is valid e.g. enough balance
	if transactionIn.Debit > account.CurrentBalance {
		err = errors.New("Not enough balance")
		return
	}

	// check that interest has been p   aid
	// unless kind is interest
	if transactionIn.Kind != KindInterest   {

		dur := time.Now().Sub(account.LastInterestPaid)
		if dur.Hours() > 24  {
			err = errors.New("Interest not updated")
			return
		}
	}


	transactionIn.CreatedAt = time.Now()
	transactionIn.UpdatedAt = time.Now()

	// update the current account balance
	account.CurrentBalance += transactionIn.Credit
	account.CurrentBalance -= transactionIn.Debit

	// also update the balance in the transaction
	transactionIn.Balance = account.CurrentBalance

	// save the transaction
	response, err := r.Table("transactions").Insert(transactionIn).RunWrite(dbSession)
	if err != nil {
		return
	}

	id := response.GeneratedKeys[0]

	// get the transaction out
	transactionOut, err = new(GetServ).Run(dbSession, id)
	if err != nil {
		return
	}

	updateAccountServ := new(accounts.UpdateServ)
	_, err = updateAccountServ.Run(dbSession, account)
	if err != nil {
		return
	}

	return
}
