package accounts_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/lib/matchers"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
)

var _ = Describe("UpdateBalanceServ", func() {

	var (
		service    accounts.UpdateBalanceServ
		createServ accounts.CreateServ
		getServ    accounts.GetServ
		account    models.Account
	)

	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func() {
		// create an empty account for testing
		account = models.Account{
			CurrentBalance:   100,
			LastInterestPaid: time.Now().AddDate(-1, 0, 0),
		}
		_, err = createServ.Run(dbSession, &account)
		if err != nil {
			fmt.Println(err)
		}

	})

	It("Saves the account", func() {
		service.Run(dbSession, &account)
		account2, _ := getServ.Run(dbSession, account.Id)
		Expect(account2.CurrentBalance).To(Equal(103.5))
	})

	It("Updates the current balance", func() {
		service.Run(dbSession, &account)
		Expect(account.CurrentBalance).To(Equal(103.5))
	})

	It("Updates the last interest paid", func() {
		service.Run(dbSession, &account)
		Expect(account.LastInterestPaid).To(matchers.BeWithin(time.Now()))
	})

	It("Doesnt update the balance twice", func() {
		service.Run(dbSession, &account)
		service.Run(dbSession, &account)
		Expect(account.CurrentBalance).To(Equal(103.5))
	})

})
