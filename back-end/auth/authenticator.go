package auth

import "fmt"

const (
	user     = "scoir"
	password = "scoir"
	token    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
)

// this probably will have some kind of config to connect an LDAP for example
func NewAuthenticatorProvider() AuthenticatorProvider {
	return authenticatorProvider{}
}

// this might seen as an overkill, but in "real life" there's a provider/driver that actually do this task,
// and it shall be injected and wrapped
type AuthenticatorProvider interface {
	AuthenticateWithLogin(login UserLogin) (*UserJTW, error)
}

type authenticatorProvider struct {
}

func (a authenticatorProvider) AuthenticateWithLogin(login UserLogin) (*UserJTW, error) {
	if login.User == user && login.Password == password {
		return &UserJTW{
			Token: token,
		}, nil
	}
	return nil, fmt.Errorf("incorrect user and password")
}
