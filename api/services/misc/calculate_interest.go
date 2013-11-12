package misc

import ()

type CalculateInterestServ struct {
}

// Calculate the interest for the given period
// at the given rate
// for the given amount
// annualRate e.g. 7.5 (%)
func (serv *CalculateInterestServ) Run(principal float64, days float64, annualRate float64) (interest float64) {
  interest = principal * days * annualRate / 100 / 365
  return
}



// // Calculate the interest for the given period
// // at the given rate
// // for the given amount
// // annualRate e.g. 7.5 (%)
// func CalculateInterest(principal float64, days float64, annualRate float64) (interest float64) {
// 	interest = principal * days * annualRate / 100 / 365
// 	return
// }
