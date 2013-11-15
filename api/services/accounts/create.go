package accounts

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"time"
)

type CreateServ struct {
}

func (serv *CreateServ) Run(dbSession *r.Session, accountIn models.Account) (accountOut models.Account, err error) {

	if accountIn.Id != "" {
		err = errors.New("Account already has an id")
		return
	}

	accountIn.CreatedAt = time.Now()
	accountIn.UpdatedAt = time.Now()

	response, err := r.Table("accounts").Insert(accountIn).RunWrite(dbSession)
	if err != nil {
		return
	}

	id := response.GeneratedKeys[0]

	accountOut, err = new(GetServ).Run(dbSession, id)

	return
}
