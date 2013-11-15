package accounts_test

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
)

var _ = Describe("GetServ", func() {

	var (
		service   accounts.GetServ
		accountId string
		createAccountServ accounts.CreateServ
	)

	dbSession, _ := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	BeforeEach(func () {
		accountIn := new(models.Account)
		accountIn.Name = "X"
		accountOut, err := createAccountServ.Run(dbSession, *accountIn)
		if err != nil {
			fmt.Println("Account not created")
		}
		accountId = accountOut.Id
	})

	It("Saved the account", func () {
		Expect(accountId).NotTo(BeEmpty())
	})

	It("Gets the account", func() {
		account, err := service.Run(dbSession, accountId)
		Expect(err).To(BeNil())
		Expect(account.Name).To(Equal("X"))
	})

})
