package main

import (
	"fmt"
	"slices"
	"vaultwarden-entra-sync/model"

	"github.com/go-resty/resty/v2"
)

func getUsers() []string {
	client := resty.New()

	var token model.Token

	resp, err := client.R().
		SetFormData(map[string]string{
			"client_id":     cfg.ClientId,
			"scope":         "https://graph.microsoft.com/.default",
			"client_secret": cfg.ClientSecret,
			"grant_type":    "client_credentials",
		}).
		SetResult(&token).
		Post(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", cfg.TenantId))

	if err != nil {
		panic(err)
	}
	if resp.IsError() {
		panic(resp.String())
	}

	client.SetAuthToken(token.AccessToken)

	var mail []string
	for _, g := range cfg.Groups {
		g := g
		mail = append(mail, getGroupMembers(client, g)...)
	}

	return slices.Compact(mail)
}

func getGroupMembers(c *resty.Client, id string) []string {
	var mail []string
	var res model.GroupMembersResponse
	var nextLink *string
	c.R().
		SetQueryParams(map[string]string{
			"$select": "mail",
			"$top":    "100",
		}).
		SetResult(&res).
		Get(fmt.Sprintf("https://graph.microsoft.com/v1.0/groups/%s/members", id))

	for _, m := range res.Value {
		m := m
		mail = append(mail, m.Mail)
	}

	nextLink = res.OdataNextLink

	for nextLink != nil {
		var res model.GroupMembersResponse
		c.R().
			SetResult(&res).
			Get(*nextLink)

		for _, m := range res.Value {
			m := m
			mail = append(mail, m.Mail)
		}

		nextLink = res.OdataNextLink
	}

	return mail
}
