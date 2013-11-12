package accounts

import (
  "github.com/sporto/kic/api/models"
  "github.com/sporto/kic/api/services/misc"
  "time"
)

type CalculateInterestToPayServ struct{
}

func (serv *CalculateInterestToPayServ) Run(account models.Account) (interest float64, err error) {
	dur, err := new(CalculateInterestDurationToPayServ).Run(account, time.Now())
	if err != nil {
		return
	}

	interest, err = new(misc.CalculateInterestServ).Run(account.CurrentBalance, dur, 3.5)
	return
}
