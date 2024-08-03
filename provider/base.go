package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type base struct {
	ctx          context.Context
	displayName  string
	clientId     string
	clientSecret string
	redirectUrl  string
	authUrl      string
	tokenUrl     string
	userApiUrl   string
	scopes       []string
}

func (p *base) Context() context.Context {
	return p.ctx
}

func (p *base) SetContext(ctx context.Context) {
	p.ctx = ctx
}

func (p *base) DisplayName() string {
	return string(p.displayName)
}

func (p *base) ClientId() string {
	return p.clientId
}

func (p *base) SetClientId(clientId string) {
	p.clientId = clientId
}

func (p *base) ClientSecret() string {
	return p.clientSecret
}

func (p *base) SetClientSecret(clientSecret string) {
	p.clientSecret = clientSecret
}

func (p *base) RedirectUrl() string {
	return p.redirectUrl
}

func (p *base) SetRedirectUrl(redirectUrl string) {
	p.redirectUrl = redirectUrl
}

func (p *base) AuthUrl() string {
	return p.authUrl
}

func (p *base) SetAuthUrl(authUrl string) {
	p.authUrl = authUrl
}

func (p *base) TokenUrl() string {
	return p.tokenUrl
}

func (p *base) SetTokenUrl(tokenUrl string) {
	p.tokenUrl = tokenUrl
}

func (p *base) UserApiUrl() string {
	return p.userApiUrl
}

func (p *base) SetUserApiUrl(userApiUrl string) {
	p.userApiUrl = userApiUrl
}

func (p *base) Scopes() []string {
	return p.scopes
}

func (p *base) SetScopes(scopes []string) {
	p.scopes = scopes
}

func (p *base) Client(token *oauth2.Token) *http.Client {
	return p.oauth2Config().Client(p.ctx, token)
}

func (p *base) FetchToken(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.oauth2Config().Exchange(p.ctx, code, opts...)
}

func (p *base) FetchRawUserData(token *oauth2.Token) (data []byte, err error) {
	req, err := http.NewRequestWithContext(p.ctx, "GET", p.userApiUrl, nil)
	if err != nil {
		return nil, err
	}

	client := p.Client(token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf(
			"failed to fetch OAuth2 user profile via %s (%d):\n%s",
			p.userApiUrl,
			res.StatusCode,
			string(result),
		)
	}

	return result, nil
}

func (p *base) oauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     p.clientId,
		ClientSecret: p.clientSecret,
		RedirectURL:  p.redirectUrl,
		Scopes:       p.scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  p.authUrl,
			TokenURL: p.tokenUrl,
		},
	}
}
