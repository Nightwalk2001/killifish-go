package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"killifish/config"

	"time"

	"killifish/docs"
)

type Claims struct {
	Name      string
	IsManager bool
	Exp       int64
}

func (c Claims) Valid() error {
	e := new(jwt.ValidationError)
	n := time.Now().Unix()
	if c.Exp < n {
		e.Errors |= jwt.ValidationErrorExpired
	}
	return e
}

func Issue(p *docs.Person) string {
	n := time.Now().Add(time.Hour * 24 * 15).Unix()
	c := Claims{
		Name:      p.Name,
		IsManager: p.IsManager,
		Exp:       n,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, _ := token.SignedString(config.Secret)
	return s
}

func (c Claims) Verify() bool {
	return c.Exp < time.Now().Unix()
}
