package transactions_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"github.com/sporto/kic/api/services/transactions"
	"time"
)

var _ = Describe("CreateServ", func() {

	var (
		createAccountServ accounts.CreateServ
		updateAccountServ accounts.UpdateServ
		service           transactions.CreateServ
		account           models.Account
		transaction       *models.Transaction
	)

	// global setup
	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func() {
		accountIn := *&models.Account{CurrentBalance: 100, LastInterestPaid: time.Now() }
		account, err = createAccountServ.Run(dbSession, accountIn)
		// fmt.Println("Account created ", account)
		if err != nil {
			fmt.Println("Account not created because ", err)
		}
		transaction = &models.Transaction{AccountId: account.Id, Debit: 50, Credit: 100}
	})

	///////////////////////////////////////////////////////////

	It("Saves the account", func () {
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("Saves the transaction", func() {
		transaction, err := service.Run(dbSession, *transaction)
		Expect(err).To(BeNil())
		Expect(transaction.Id).NotTo(BeEmpty())
	})

	It("Updates the account balance", func() {
		_, err = service.Run(dbSession, *transaction)
		Expect(err).To(BeNil())
		getAccountServ := new(accounts.GetServ)
		account, err = getAccountServ.Run(dbSession, account.Id)
		Expect(account.CurrentBalance).To(Equal(150.0))
	})

	It("fails when no account id provided", func() {
		transaction.AccountId = ""
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't exist", func() {
		transaction.AccountId = "XYZ"
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the transaction is already saved", func() {
		transaction.Id = "aaaa"
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if no credit or debit provided", func() {
		transaction.Credit = 0
		transaction.Debit = 0
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't have enough balance", func() {
		transaction.Debit = 150
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account interests has not been updated", func () {
		account.LastInterestPaid = time.Now().AddDate(0, 0, -2) // two days
		_, err = updateAccountServ.Run(dbSession, account)
		Expect(err).To(BeNil())
		_, err = service.Run(dbSession, *transaction)
		Expect(err).NotTo(BeNil())
	})
})
