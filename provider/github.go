package provider

import (
	"context"
	"encoding/json"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var _ Provider = (*Github)(nil)

const NameGithub = "github"

type Github struct {
	*base
}

func NewGithub() *Github {
	return &Github{
		base: &base{
			ctx:        context.Background(),
			scopes:     []string{"read:user", "user:email"},
			authUrl:    github.Endpoint.AuthURL,
			tokenUrl:   github.Endpoint.TokenURL,
			userApiUrl: "https://api.github.com/user",
		},
	}
}

func (p *Github) FetchUser(token *oauth2.Token) (*AuthUser, error) {
	data, err := p.FetchRawUserData(token)
	if err != nil {
		return nil, err
	}

	rawUser := map[string]any{}
	if err := json.Unmarshal(data, &rawUser); err != nil {
		return nil, err
	}

	extracted := struct {
		Id        int    `json:"id"`
		Username  string `json:"login"`
		Email     string `json:"email"`
		AvatarUrl string `json:"avatar_url"`
	}{}

	if err := json.Unmarshal(data, &extracted); err != nil {
		return nil, err
	}

	user := &AuthUser{
		Id:        strconv.Itoa(extracted.Id),
		Username:  extracted.Username,
		Email:     extracted.Email,
		AvatarUrl: extracted.AvatarUrl,
	}
	return user, nil
}
