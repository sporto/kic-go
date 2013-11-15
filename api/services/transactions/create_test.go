package transactions_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"github.com/sporto/kic/api/services/transactions"
)

var _ = Describe("CreateServ", func() {

	var (
		createAccountServ accounts.CreateServ
		service           transactions.CreateServ
		account           models.Account
	)

	// global setup
	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func() {
		accountIn := *&models.Account{CurrentBalance: 100}
		account, err = createAccountServ.Run(dbSession, accountIn)
		// fmt.Println("Account created ", account)
		if err != nil {
			fmt.Println("Account not created")
		}
	})

	///////////////////////////////////////////////////////////

	It("Saves the account", func () {
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("Saves the transaction", func() {
		transaction := models.Transaction{AccountId: account.Id, Credit: 100}
		transaction, err := service.Run(dbSession, transaction)
		Expect(err).To(BeNil())
		Expect(transaction.Id).NotTo(BeEmpty())
	})

	It("fails when no account id provided", func() {
		transaction := models.Transaction{Credit: 100}
		_, err = service.Run(dbSession, transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't exist", func() {
		transaction := models.Transaction{AccountId: "ZZZZ", Credit: 100}
		_, err = service.Run(dbSession, transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the transaction is already saved", func() {
		transaction := models.Transaction{Id: "aaaaa", AccountId: account.Id, Credit: 100}
		_, err = service.Run(dbSession, transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if no credit or debit provided", func() {
		transaction := models.Transaction{AccountId: account.Id}
		_, err = service.Run(dbSession, transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't have enough balance", func() {
		transaction := models.Transaction{AccountId: account.Id, Debit: 150}
		_, err = service.Run(dbSession, transaction)
		Expect(err).NotTo(BeNil())
	})

	It("Updates the account balance", func() {
		transaction := models.Transaction{AccountId: account.Id, Debit: 50, Credit: 100}
		_, err = service.Run(dbSession, transaction)
		Expect(err).To(BeNil())
		getAccountServ := new(accounts.GetServ)
		account, err = getAccountServ.Run(dbSession, account.Id)
		Expect(account.CurrentBalance).To(Equal(150.0))
	})
})
