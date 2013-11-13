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
	)

	dbSession, err := r.Connect(map[string]interface{}{
		"address":  "localhost:28015",
		"database": "kic_test",
	})

	// create an account
	createAccountServ := *new(accounts.CreateServ)
	accountId, err = createAccountServ.Run(dbSession, &*new(models.Account))
	if err != nil {
		fmt.Println("Account not created")
	}

	It("Gets the account", func() {
		account, err := service.Run(dbSession, accountId)
		Expect(err).To(BeNil())
		Expect(account.Id).To(Equal(accountId))
	})

})
