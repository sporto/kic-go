package accounts_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
)

var _ = Describe("UpdateServ", func() {

	var (
		service    accounts.UpdateServ
		createServ accounts.CreateServ
		account  models.Account
	)

	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func () {
		// create an account for testing
		accountIn := &models.Account{}
		account, err = createServ.Run(dbSession, *accountIn)

		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("Account saved", accountIn.Id)
	})

	It("Saved the account", func () {
		Expect(account.Id).NotTo(BeEmpty())
	})

	It("Updates the account", func() {
		account.Name = "No name"
		_, err := service.Run(dbSession, account)
		Expect(err).To(BeNil())
	})

	It("fails if account doesnt have an id", func() {
		account.Id = ""
		_, err := service.Run(dbSession, account)
		Expect(err).NotTo(BeNil())
	})

})
