package accounts_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"log"
)

var _ = Describe("UpdateServ", func() {

	var (
		service    accounts.UpdateServ
		createServ accounts.CreateServ
		account    models.Account
	)

	dbSession, err := api.GetDbSession("../../../")
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	BeforeEach(func() {
		// create an account for testing
		accountIn := &models.Account{}
		account, err = createServ.Run(dbSession, *accountIn)

		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("Account saved", accountIn.Id)
	})

	It("Saved the account", func() {
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
