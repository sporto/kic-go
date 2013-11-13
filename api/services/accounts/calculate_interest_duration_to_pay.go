package accounts

import (
	"github.com/sporto/kic/api/models"
	"time"
)

type CalculateInterestDurationToPayServ struct {
}

func (serv *CalculateInterestDurationToPayServ) Run(account models.Account, to time.Time) (days time.Duration, err error) {
	zero := *new(time.Time)
	if account.LastInterestPaid.Equal(zero) {
		days = 0
	} else {
		days = to.Sub(account.LastInterestPaid)
	}
	return
}
