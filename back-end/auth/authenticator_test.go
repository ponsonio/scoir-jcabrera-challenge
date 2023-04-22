package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticatorProvider_AuthenticateWithLogin(t *testing.T) {

	prov := NewAuthenticatorProvider()

	u := UserLogin{
		User:     user,
		Password: password,
	}

	tok, err := prov.AuthenticateWithLogin(u)

	assert.Equal(t, token, tok.Token)
	assert.Nil(t, nil)

	u.User = "notanuser"

	tok, err = prov.AuthenticateWithLogin(u)

	assert.Nil(t, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "incorrect user and password", err.Error())

	u.Password = "badpasswod"

	tok, err = prov.AuthenticateWithLogin(u)

	assert.Nil(t, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "incorrect user and password", err.Error())

}
