# Auth
a simple way to authenticate with oauth2 providers

# TODO
- [ ] jwt token support
- [ ] Add more providers

## Supported Providers
- [x] Github
- [x] Google

## Usage
```go
// 1. Create a new auth client with the provider
conf, err := auth.NewConfig(provider.NameGithub, 
		auth.WithClientId("YOUR_CLIENT_ID"),
		auth.WithClientSecret("YOUR_CLIENT_SECRET"),
		auth.WithRedirectUrl("YOUR_REDIRECT_URL"))

// 2. Use the code to exchange for a token
token, err := conf.GetToken(code)

// 3. Use the token to get user Message
user, err := conf.GetUser(token)
```

# References
- [oauth2](golang.org/x/oauth2)
- [pocketbase/auth](https://github.com/pocketbase/pocketbase/tree/master/tools/auth)