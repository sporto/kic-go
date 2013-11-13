package accounts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
)

var _ = Describe("CalculateInterestToPayServ", func() {

	var (
		serv accounts.CalculateInterestToPayServ
	)

	It("returns 0 if no date given", func() {
		account := models.Account{}
		interest, _ := serv.Run(account)
		Expect(interest).To(Equal(0.0))
	})

	It("returns the correct intersest to pay", func() {
		now := time.Now()
		d := now.AddDate(-1, 0, 0)
		account := models.Account{CurrentBalance: 100, LastInterestPaid: d}
		interest, _ := serv.Run(account)
		Expect(interest).To(Equal(3.5))
	})

})
