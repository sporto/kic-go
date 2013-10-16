package controllers

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/api/models"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net/http"
)

type Accounts struct {
	DbSession *r.Session
}

// func (c *Accounts) Options(ctx context.Context) error {
// 	log.Println("Options for Account")
// 	return nil
// }

func (c *Accounts) ReadMany(ctx context.Context) error {
	var records []models.Account

	err := r.Table("accounts").Run(c.DbSession).All(&records)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, records)
}

func (c *Accounts) Read(id string, ctx context.Context) error {
	var record models.Account

	err := r.Table("accounts").Get(id).Run(c.DbSession).One(&record)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	if record.Id == "" {
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	return goweb.API.RespondWithData(ctx, record)
}
