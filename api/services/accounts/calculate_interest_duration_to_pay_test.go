package accounts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"time"
)

var _ = Describe("CalculateInterestDurationToPayServ", func() {

	var (
		serv accounts.CalculateInterestDurationToPayServ
	)

	It("returns 0 if no date given", func() {
		account := models.Account{}
		now := time.Now()
		dur, _ := serv.Run(account, now)
		expectedDur := time.Duration(0)
		Expect(dur).To(Equal(expectedDur))
	})

	It("returns the right number", func() {
		now := time.Now()
		ti := now.AddDate(0, 0, -4)
		account := models.Account{LastInterestPaid: ti}
		dur, _ := serv.Run(account, now)
		expectedDur := now.Sub(ti)
		Expect(dur).To(Equal(expectedDur))
	})

})
