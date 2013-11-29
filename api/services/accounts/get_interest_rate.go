package accounts

import (
  "github.com/sporto/kic/api/models"
)

type GetInterestRateServ struct {
}

// calculate the yearly interest rate
func (serv *GetInterestRateServ) Run(account models.Account) (rate float64, err error) {
	rate = 50
	return
}
