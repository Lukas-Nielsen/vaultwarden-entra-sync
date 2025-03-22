package main

import (
	"fmt"
	"log"
	"slices"
	"vaultwarden-entra-sync/model"

	"github.com/Lukas-Nielsen/go-vaultwarden-admin"
	"github.com/caarlos0/env/v11"
	"github.com/go-co-op/gocron/v2"
)

var cfg model.Config

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	s, err := gocron.NewScheduler()

	if err != nil {
		log.Fatalln(err)
	}

	_, err = s.NewJob(
		gocron.CronJob(
			cfg.Cron, false,
		),
		gocron.NewTask(
			func() {
				run()
			},
		),
	)
	if err != nil {
		panic(err)
	}

	s.Start()

	select {}
}

func run() {
	toDisable := []string{}
	toInvite := []string{}

	mails := getUsers()

	c, err := vaultwarden.NewClient(cfg.BaseUrl, cfg.AdminToken, false)

	if err != nil {
		panic(err)
	}

	users, err := c.UsersGetAll()

	if err != nil {
		panic(err)
	}

	for _, m := range mails {
		m := m
		if !slices.ContainsFunc(users, func(e vaultwarden.User) bool {
			return e.Email == m
		}) {
			toInvite = append(toInvite, m)
		}
	}

	for _, u := range users {
		u := u
		if !slices.Contains(mails, u.Email) && !slices.Contains(cfg.Whitelist, u.Email) {
			toDisable = append(toDisable, u.ID)
		}
	}

	for _, e := range toInvite {
		e := e
		_, err := c.Invite(e)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("user with mail '%s' invited\n", e)
		}
	}

	for _, e := range toDisable {
		e := e
		err := c.UsersDisable(e)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("user with id '%s' disabled\n", e)
		}
	}
}
