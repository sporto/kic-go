package misc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMisc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Misc Suite")
}
