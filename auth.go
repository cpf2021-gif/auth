package auth

import (
	"github.com/cpf2021-gif/auth/provider"
	"golang.org/x/oauth2"
)

type Config struct {
	provider provider.Provider
}

func NewConfig(providerName string, opts ...ProviderOption) (*Config, error) {
	provider, err := provider.NewProvider(providerName)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(provider)
	}

	return &Config{provider: provider}, nil
}

func (c *Config) GetToken(code string) (*oauth2.Token, error) {
	token, err := c.provider.FetchToken(code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (c *Config) GetUser(token *oauth2.Token) (*provider.AuthUser, error) {
	user, err := c.provider.FetchUser(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type ProviderOption func(provider.Provider)

func WithUserApiUrl(userApiUrl string) ProviderOption {
	return func(p provider.Provider) {
		p.SetUserApiUrl(userApiUrl)
	}
}

func WithRedirectUrl(redirectUrl string) ProviderOption {
	return func(p provider.Provider) {
		p.SetRedirectUrl(redirectUrl)
	}
}

func WithClientId(clientId string) ProviderOption {
	return func(p provider.Provider) {
		p.SetClientId(clientId)
	}
}

func WithClientSecret(clientSecret string) ProviderOption {
	return func(p provider.Provider) {
		p.SetClientSecret(clientSecret)
	}
}
