package auth

import (
	"errors"

	"github.com/cpf2021-gif/auth/provider"
	"golang.org/x/oauth2"
)

var providers map[string]provider.Provider = make(map[string]provider.Provider)

func RegisterProviders(ps ...provider.Provider) {
	for _, p := range ps {
		providers[p.DisplayName()] = p
	}
}

func getProvider(providerName string) (provider.Provider, error) {
	p, ok := providers[providerName]
	if !ok {
		return nil, errors.New("provider not found")
	}

	return p, nil
}

func GetToken(providerName, code string) (*oauth2.Token, error) {
	p, err := getProvider(providerName)
	if err != nil {
		return nil, err
	}

	token, err := p.FetchToken(code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUser(providerName string, token *oauth2.Token) (*provider.AuthUser, error) {
	p, err := getProvider(providerName)
	if err != nil {
		return nil, err
	}

	user, err := p.FetchUser(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
