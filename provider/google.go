package provider

import (
	"context"
	"fmt"
	"io"

	"github.com/imroc/req/v3"
	"golang.org/x/oauth2"
)

var _ Provider = (*Google)(nil)

const NameGoogle string = "google"

type Google struct {
	*base
}

func NewGoogle(opts ...ProviderOption) *Google {
	p := &Google{
		base: &base{
			ctx:         context.Background(),
			displayName: NameGoogle,
			scopes: []string{
				"openid",
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			authUrl:    "https://accounts.google.com/o/oauth2/v2/auth",
			tokenUrl:   "https://oauth2.googleapis.com/token",
			userApiUrl: "https://www.googleapis.com/oauth2/v2/userinfo",
		},
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Google) FetchToken(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	client := req.C()
	token := &oauth2.Token{}

	payload := fmt.Sprintf("client_id=%s&client_secret=%s&redirect_uri=%s&code=%s&grant_type=authorization_code", p.ClientId(), p.ClientSecret(), p.RedirectUrl(), code)

	res, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBodyString(payload).
		SetSuccessResult(token).
		Post(p.TokenUrl())
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
			"failed to fetch OAuth2 token via %s (%d):\n%s",
			p.TokenUrl(),
			res.StatusCode,
			string(result),
		)
	}

	return token, nil
}

func (p *Google) FetchUser(token *oauth2.Token) (*AuthUser, error) {
	client := req.C()

	extracted := struct {
		Id        string `json:"id"`
		Username  string `json:"name"`
		Email     string `json:"email"`
		AvatarUrl string `json:"picture"`
	}{}

	res, err := client.R().
		SetQueryParam("access_token", token.AccessToken).
		SetSuccessResult(&extracted).
		Get(p.userApiUrl)
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

	user := &AuthUser{
		Id:        extracted.Id,
		Username:  extracted.Username,
		Email:     extracted.Email,
		AvatarUrl: extracted.AvatarUrl,
	}

	return user, nil
}
