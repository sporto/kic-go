package controllers

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
	"github.com/sporto/kic/api/services/account_balances"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net/http"
)

type Accounts struct {
	DbSession *r.Session
}

func (c *Accounts) Path() string {
	return "api/accounts"
}

func (c *Accounts) ReadMany(ctx context.Context) (err error) {
	var accounts []models.Account

	rows, err := r.Table("accounts").OrderBy(r.Asc("CreatedAt")).Run(c.DbSession)
	rows.Scan(&accounts)

	for rows.Next() {
		var account models.Account

		err = rows.Scan(&account)
		if err != nil {
			return
		}

		accounts = append(accounts, account)
	}

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, accounts)
}

func (c *Accounts) Read(id string, ctx context.Context) (err error) {

	getServ := &accounts.GetServ{}
	account, err := getServ.Run(c.DbSession, id)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	if account.Id == "" {
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	return goweb.API.RespondWithData(ctx, account)
}

func (c *Accounts) Adjust(ctx context.Context) (err error) {
	id := ctx.PathParams().Get("id").Str()

	getServ := &accounts.GetServ{}
	account, err := getServ.Run(c.DbSession, id)
	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusNotFound, err.Error())
	}

	adjustServ := &account_balances.AdjustServ{}
	account, _, err = adjustServ.Run(c.DbSession, account)
	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return goweb.API.RespondWithData(ctx, account)
}
