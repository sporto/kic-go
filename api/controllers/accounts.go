package controllers

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/accounts"
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

func (c *Accounts) ReadMany(ctx context.Context) error {
	var accounts []models.Account

	rows, err := r.Table("accounts").OrderBy(r.Asc("CreatedAt")).Run(c.DbSession)
	rows.Scan(&accounts)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, accounts)
}

func (c *Accounts) Read(id string, ctx context.Context) error {
	var record models.Account

	getServ := &accounts.GetServ{}
	account, err := getServ.Run(c.DbSession, id)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	if record.Id == "" {
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	return goweb.API.RespondWithData(ctx, account)
}
