package misc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"fmt"
  "time"
	"github.com/sporto/kic/api/services/misc"
)

var _ = Describe("CalculateServ", func() {

	var (
		service misc.CalculateInterestServ
	)

	It("Calculates the right interest", func () {
    d := time.Duration(365*24*time.Hour)
		i, _ := service.Run(100, d, 3.0)
		fmt.Println(i)
		Expect(i).To(Equal(3.0))
	})

	It("Calculates the right interest", func () {
    d := time.Duration(10*24*time.Hour)
		i, _ := service.Run(100, d, 3.0)
		Expect(math.Floor(i*1000)).To(Equal(82.0))
	})

})