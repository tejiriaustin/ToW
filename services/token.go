package services

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/tejiriaustin/ToW/env"
)

type (
	TokenGenerator struct {
		conf *env.Config
	}

	TokenProvider interface {
		GenerateToken(conf *env.Config, content any) (string, error)
	}

	Claims struct {
		Exp           time.Time
		Authorization bool
		jwt.StandardClaims
		Content any
	}
)

func NewTokenProvider() TokenProvider {
	return &TokenGenerator{}
}

func (p *TokenGenerator) GenerateToken(conf *env.Config, content any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Exp:           time.Now().Add(3600 * time.Minute),
		Authorization: true,
		Content:       content,
	})

	pkey := conf.GetAsBytes(env.JwtSecret)
	tokenString, err := token.SignedString(pkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
