package auth

type AuthenticationService interface {
	AuthenticateWithLogin(login UserLogin) (*UserJTW, error)
}

type authenticationService struct {
	authenticatorProvider *AuthenticatorProvider
}

func NewAuthenticationService(authenticatorProvider *AuthenticatorProvider) AuthenticationService {
	ret := &authenticationService{
		authenticatorProvider: authenticatorProvider,
	}
	return ret
}

func (srv *authenticationService) AuthenticateWithLogin(login UserLogin) (*UserJTW, error) {
	p := *srv.authenticatorProvider
	return p.AuthenticateWithLogin(login)
}
