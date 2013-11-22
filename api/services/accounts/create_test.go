package accounts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"log"
	// "fmt"
)

var _ = Describe("CreateServ", func() {

	var (
		service accounts.CreateServ
	)

	dbSession, err := api.GetDbSession("../../../")
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	It("Saves the account", func() {
		accountIn := &models.Account{Name: "Sam"}
		accountOut, err := service.Run(dbSession, *accountIn)
		Expect(err).To(BeNil())
		Expect(accountOut.Id).NotTo(BeEmpty())
	})

	It("fails if the account is already saved", func() {
		account := &models.Account{Id: "aaaaa"}
		_, err = service.Run(dbSession, *account)
		Expect(err).NotTo(BeNil())
	})

})
