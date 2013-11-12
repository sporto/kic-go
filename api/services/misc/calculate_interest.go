package misc

import (
  "time"
)

type CalculateInterestServ struct {
}

// Calculate the interest for the given period
// at the given rate
// for the given amount
// annualRate e.g. 7.5 (%)
func (serv *CalculateInterestServ) Run(principal float64, dur time.Duration, annualRate float64) (interest float64, err error) {
	days := float64(dur / 24 / time.Hour)
	interest = principal * days * annualRate / 100 / 365
	return
}