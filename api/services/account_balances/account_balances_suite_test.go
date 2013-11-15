package account_balances_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAccount_balances(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account_balances Suite")
}
