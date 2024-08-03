# Auth
a simple way to authenticate with oauth2 providers

# TODO
- [ ] jwt token support
- [ ] Add more providers

## Supported Providers
- [x] Github
- [x] Google

## Example
```go
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
```

# References
- [oauth2](golang.org/x/oauth2)
- [pocketbase/auth](https://github.com/pocketbase/pocketbase/tree/master/tools/auth)