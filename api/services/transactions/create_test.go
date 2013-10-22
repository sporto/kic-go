package transactions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/sporto/kic/api/models"
)

var _ = Describe("Create", func() {
	It("fails when no account id provided", func() {
		Expect(1).To(Equal(1))
	})

	It("fails if the account id doesn't exist")

	It("fails if the transaction is already saved")

	It("fails if the account doesn't have enough balance")

	It("Saves the transaction")

	It("Updates the account balance")
})
