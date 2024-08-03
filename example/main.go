package main

import (
	"fmt"

	"github.com/cpf2021-gif/auth"
	"github.com/cpf2021-gif/auth/provider"
)

func main() {
	// 1. Register providers
	auth.RegisterProviders(
		provider.NewGithub(
			provider.WithClientId("client_id"),
			provider.WithClientSecret("client_secret"),
			provider.WithRedirectUrl("redirect_url"),
		),
		provider.NewGoogle(
			provider.WithClientId("client_id"),
			provider.WithClientSecret("client_secret"),
			provider.WithRedirectUrl("redirect_url"),
		),
	)

	// 2. Get token
	token, err := auth.GetToken("github", "code")
	if err != nil {
		fmt.Println(err)
	}

	// 3. Get user
	user, err := auth.GetUser("github", token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
