package accounts_test

import (
	r "github.com/dancannon/gorethink"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sporto/kic/api"
	"testing"
	"log"
	"errors"
	"os"
)

var dbSession *r.Session

func TestAccounts(t *testing.T) {
	RegisterFailHandler(Fail)

	os.Setenv("ENV", "test")

	err := errors.New("")
	dbSession, err = api.StartDb("../../../")
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	RunSpecs(t, "Accounts Suite")
}
