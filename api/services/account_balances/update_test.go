package account_balances_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/lib/matchers"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"github.com/sporto/kic/api/services/account_balances"
	"time"
)

var _ = Describe("UpdateServ", func() {

	var (
		service    account_balances.UpdateServ
		createServ accounts.CreateServ
		// getServ    accounts.GetServ
		account    models.Account
	)

	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func() {
		// create an empty accountIn for testing
		accountIn := &models.Account{
			CurrentBalance:   100,
			LastInterestPaid: time.Now().AddDate(-1, 0, 0),
		}
		account, err = createServ.Run(dbSession, *accountIn)
		if err != nil {
			fmt.Println(err)
		}
	})

	It("Saved the test account", func () {
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("Returns the account", func() {
		accountOut, _, err := service.Run(dbSession, account)
		Expect(err).To(BeNil())
		Expect(accountOut.Id).NotTo(BeEmpty())
	})

	It("Updates the current balance", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		Expect(accountOut.CurrentBalance).To(Equal(103.5))
	})

	It("Updates the last interest paid", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		Expect(accountOut.LastInterestPaid).To(matchers.BeWithin(time.Now()))
	})

	It("Doesnt update the balance twice", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		accountOut, _, _ = service.Run(dbSession, account)
		Expect(accountOut.CurrentBalance).To(Equal(103.5))
	})

	It("Creates a transaction", func () {
		_, transaction, _ := service.Run(dbSession, account)
		Expect(transaction.Credit).To(Equal(3.5))
		Expect(transaction.Id).NotTo(BeEmpty())
	})

})
