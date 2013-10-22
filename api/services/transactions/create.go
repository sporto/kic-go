package transactions

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/api/models"
)

func Create(dbSession *r.Session, transaction models.Transaction) (r.WriteResponse, error) {
	// fail if transaction is already saved
	// check that the account exist
	// check that the transaction is valid e.g. enough balance
	// save the transaction
	// update the current account balance

	var record r.WriteResponse
	err := r.Table("transactions").Insert(transaction).Run(dbSession).One(&record)

	return record, err
}
