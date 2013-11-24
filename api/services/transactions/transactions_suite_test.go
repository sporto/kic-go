package transactions_test

import (
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api"
	"testing"
	"log"
	"errors"
)

var dbSession *r.Session

func TestTransactions(t *testing.T) {
	RegisterFailHandler(Fail)

	err := errors.New("")
	dbSession, err = api.StartDb("../../../")
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	RunSpecs(t, "Transactions Suite")
}
