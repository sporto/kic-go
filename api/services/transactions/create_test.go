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
		accountId         string
		account           models.Account
	)

	// global setup
	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	// create an account for the test transactions
	// createAccountServ = *new(accounts.CreateServ)
	// account := *&models.Account{CurrentBalance: 100}
	// accountId, err = createAccountServ.Run(dbSession, &account)
	// if err != nil {
	// 	fmt.Println("Account not created")
	// }

	BeforeEach(func() {
		account = *&models.Account{CurrentBalance: 100}
		accountId, err = createAccountServ.Run(dbSession, &account)
		if err != nil {
			fmt.Println("Account not created")
		}
	})

	///////////////////////////////////////////////////////////

	It("Saves the transaction", func() {
		transaction := models.Transaction{AccountId: accountId, Credit: 100}
		id, err := service.Run(dbSession, &transaction)
		Expect(err).To(BeNil())
		Expect(id).NotTo(BeEmpty())
	})

	It("Adds the id to the transaction", func() {
		transaction := models.Transaction{AccountId: accountId, Credit: 100}
		_, err := service.Run(dbSession, &transaction)
		Expect(err).To(BeNil())
		Expect(transaction.Id).NotTo(BeEmpty())
	})

	It("fails when no account id provided", func() {
		transaction := models.Transaction{Credit: 100}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't exist", func() {
		transaction := models.Transaction{AccountId: "ZZZZ", Credit: 100}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the transaction is already saved", func() {
		transaction := models.Transaction{Id: "aaaaa", AccountId: accountId, Credit: 100}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if no credit or debit provided", func() {
		transaction := models.Transaction{AccountId: accountId}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).NotTo(BeNil())
	})

	It("fails if the account doesn't have enough balance", func() {
		transaction := models.Transaction{AccountId: accountId, Debit: 150}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).NotTo(BeNil())
	})

	It("Updates the account balance", func() {
		transaction := models.Transaction{AccountId: accountId, Debit: 50, Credit: 100}
		_, err = service.Run(dbSession, &transaction)
		Expect(err).To(BeNil())
		getAccountServ := new(accounts.GetServ)
		account, err = getAccountServ.Run(dbSession, accountId)
		Expect(account.CurrentBalance).To(Equal(150.0))
	})
})
