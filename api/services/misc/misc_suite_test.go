package misc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"os"
)

func TestMisc(t *testing.T) {
	RegisterFailHandler(Fail)

	os.Setenv("ENV", "test")

	RunSpecs(t, "Misc Suite")
}
