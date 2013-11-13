package accounts

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"time"
)

type UpdateServ struct {
}

func (serv *UpdateServ) Run(dbSession *r.Session, account models.Account) (err error) {

	if account.Id == "" {
		err = errors.New("Invalid id")
		return
	}

	account.UpdatedAt = time.Now()

	_, err = r.Table("accounts").Get(account.Id).Update(account).RunRow(dbSession)

	if err != nil {
		return
	}

	return
}
