package transactions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTransactions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transactions Suite")
}
