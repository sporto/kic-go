package account_balances_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/lib/matchers"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"github.com/sporto/kic/api/services/account_balances"
	"time"
)

var _ = Describe("AdjustServ", func() {

	var (
		service     account_balances.AdjustServ
		createServ  accounts.CreateServ
		getRateServ accounts.GetInterestRateServ
		account     models.Account
		err         error
		rate        float64
	)

	BeforeEach(func() {
		// create an empty accountIn for testing
		accountIn := &models.Account{
			CurrentBalance:   100,
			LastInterestPaid: time.Now().AddDate(-1, 0, 0),
		}
		// create the account and get a ref
		account, err = createServ.Run(dbSession, *accountIn)
		if err != nil {
			fmt.Println(err)
		}
		rate, err = getRateServ.Run(account)
	})

	It("saved the test account", func () {
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("returns the account", func() {
		accountOut, _, err := service.Run(dbSession, account)
		Expect(err).To(BeNil())
		Expect(accountOut.Id).NotTo(BeEmpty())
	})

	It("does nothing if duration is less than one day", func () {
		account.LastInterestPaid = time.Now().Add(-time.Duration(time.Hour * 10))
		accountOut, transaction, _ := service.Run(dbSession, account)
		Expect(accountOut.CurrentBalance).To(Equal(100.0))
		Expect(transaction.Id).To(BeEmpty())
	})

	It("updates the current balance", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		Expect(accountOut.CurrentBalance).To(Equal(100 + rate))
	})

	It("updates the last interest paid", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		Expect(accountOut.LastInterestPaid).To(matchers.BeWithin(time.Now()))
	})

	It("doesnt update the balance twice", func() {
		accountOut, _, _ := service.Run(dbSession, account)
		accountOut, _, _ = service.Run(dbSession, account)
		Expect(accountOut.CurrentBalance).To(Equal(100 + rate))
	})

	It("creates a transaction", func () {
		_, transaction, _ := service.Run(dbSession, account)
		Expect(transaction.Credit).To(Equal(rate))
		Expect(transaction.Id).NotTo(BeEmpty())
	})

})
