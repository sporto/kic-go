package accounts

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"errors"
	"time"
)

type CreateServ struct {
}

func (serv *CreateServ) Run(dbSession *r.Session, account models.Account) (id string, err error) {

	if account.Id != "" {
		err = errors.New("Account already has an id")
		return
	}

	account.CreatedAt = time.Now()

	response, err := r.Table("accounts").Insert(account).RunWrite(dbSession)
	if err != nil {
		return
	}
	
	id = response.GeneratedKeys[0]

	return
}
