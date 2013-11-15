package accounts

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"time"
)

type UpdateServ struct {
}

func (serv *UpdateServ) Run(dbSession *r.Session, accountIn models.Account) (accountOut models.Account, err error) {

	if accountIn.Id == "" {
		err = errors.New("Invalid id")
		return
	}

	accountIn.UpdatedAt = time.Now()

	_, err = r.Table("accounts").Get(accountIn.Id).Update(accountIn).RunRow(dbSession)
	if err != nil {
		return
	}

	accountOut, err = new(GetServ).Run(dbSession, accountIn.Id)

	return
}
