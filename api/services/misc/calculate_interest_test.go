package misc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"fmt"
	"github.com/sporto/kic/api/services/misc"
)

var _ = Describe("CalculateServ", func() {

	var (
		service misc.CalculateInterestServ
	)

	It("Calculates the right interest", func () {
		fmt.Println("xcvljsdlfjasljfkladsjkljsafklj")
		i := service.Run(100, 365, 3)
		fmt.Println(i)
		Expect(i).To(Equal(3.0))
	})

	It("Calculates the right interest", func () {
		i := service.Run(100, 10, 3)
		Expect(math.Floor(i*1000)).To(Equal(82.0))
	})

})