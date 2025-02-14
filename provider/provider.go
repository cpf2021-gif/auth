package provider

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type AuthUser struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

type Provider interface {
	Context() context.Context
	SetContext(ctx context.Context)

	DisplayName() string

	Scopes() []string
	SetScopes(scopes []string)

	ClientId() string
	SetClientId(clientId string)

	ClientSecret() string
	SetClientSecret(clientSecret string)

	RedirectUrl() string
	SetRedirectUrl(redirectUrl string)

	AuthUrl() string
	SetAuthUrl(authUrl string)

	TokenUrl() string
	SetTokenUrl(tokenUrl string)

	UserApiUrl() string
	SetUserApiUrl(userApiUrl string)

	Client(token *oauth2.Token) *http.Client
	FetchToken(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	FetchUser(token *oauth2.Token) (user *AuthUser, err error)
}

type ProviderOption func(Provider)

func WithRedirectUrl(redirectUrl string) ProviderOption {
	return func(p Provider) {
		p.SetRedirectUrl(redirectUrl)
	}
}

func WithClientId(clientId string) ProviderOption {
	return func(p Provider) {
		p.SetClientId(clientId)
	}
}

func WithClientSecret(clientSecret string) ProviderOption {
	return func(p Provider) {
		p.SetClientSecret(clientSecret)
	}
}
