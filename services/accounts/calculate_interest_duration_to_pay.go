package accounts

import (
	"github.com/sporto/kic-api/models"
	"time"
)

func CalculateInterestDurationToPay(account models.Account, to time.Time) (days time.Duration) {
	zero := *new(time.Time)
	if account.LastInterestPaid.Equal(zero) {
		days = 0
	} else {
		days = to.Sub(account.LastInterestPaid)
	}
	return
}
