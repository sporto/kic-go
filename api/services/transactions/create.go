package transactions

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"errors"
	// "fmt"
	"time"
)

type CreateServ struct {
}

func (serv *CreateServ) Run(dbSession *r.Session, transaction *models.Transaction) (id string, err error) {
	
	// fail if transaction is already saved
	if transaction.Id != "" {
		err = errors.New("Transaction Id must be nil")
		return
	}

	// fail if not account id provided
	if transaction.AccountId == "" {
		err = errors.New("Account Id cannot be nil")
		return
	}

	// fail if no credit or debit provided
	if transaction.Credit <= 0 && transaction.Debit <= 0 {
		err = errors.New("Credit or Debit must be provided")
		return
	}

	// check that the account exist
	getServ := new(accounts.GetServ)
	account, err := getServ.Run(dbSession, transaction.AccountId)
	if err != nil {
		return
	}

	// check that the transaction is valid e.g. enough balance
	if (transaction.Debit > account.CurrentBalance) {
		err = errors.New("Not enough balance")
		return
	}

	transaction.CreatedAt = time.Now()

	// save the transaction
	response, err := r.Table("transactions").Insert(transaction).RunWrite(dbSession)
	if err != nil {
		return
	}

	id = response.GeneratedKeys[0]

	transaction.Id = id

	// update the current account balance
	account.CurrentBalance += transaction.Credit
	account.CurrentBalance -= transaction.Debit
	updateAccountServ := new(accounts.UpdateServ)
	err = updateAccountServ.Run(dbSession, account)
	if err != nil {
		return
	}

	return
}
