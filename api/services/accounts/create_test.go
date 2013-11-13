package accounts_test

import (
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	// "fmt"
)

var _ = Describe("CreateServ", func() {

	var (
		service accounts.CreateServ
	)

	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	It("Saves the account", func() {
		account := models.Account{Name: "Sam"}
		id, err := service.Run(dbSession, &account)
		Expect(err).To(BeNil())
		Expect(id).NotTo(BeEmpty())
	})

	It("Adds the created id to the given account", func() {
		account := models.Account{Name: "Sam"}
		_, err := service.Run(dbSession, &account)
		Expect(err).To(BeNil())
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("fails if the account is already saved", func() {
		account := models.Account{Id: "aaaaa"}
		_, err = service.Run(dbSession, &account)
		Expect(err).NotTo(BeNil())
	})

})
