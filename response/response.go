package response

import (
	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/models"
)

func SingleAccountResponse(account *models.Account) map[string]interface{} {
	return map[string]interface{}{
		"_id":       account.ID.Hex(),
		"email":     account.Email,
		"firstName": account.FirstName,
		"lastName":  account.LastName,
		"phone":     account.Phone,
		"token":     account.Token,
		"country":   account.DOB,
	}
}

func MultipleAccountResponse(accounts []models.Account) interface{} {
	m := make([]map[string]interface{}, 0, len(accounts))
	for _, a := range accounts {
		m = append(m, SingleAccountResponse(&a))
	}
	return m
}

func GetConfigsResponse(conf *env.Config) map[string]interface{} {
	return map[string]interface{}{
		env.JwtSecret:          conf.GetAsString(env.JwtSecret),
		env.EnvPort:            conf.GetAsString(env.EnvPort),
		env.MinimumFollowSpend: conf.GetAsString(env.MinimumFollowSpend),
	}
}
