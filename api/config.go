package api

import (
	"log"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	OAuthConfig *oauth2.Config

	SessionStore sessions.Store
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	OAuthConfig = configureOAuthClient(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	cookieStore := sessions.NewCookieStore(securecookie.GenerateRandomKey(10))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:8000/callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{},
		Endpoint:     spotify.Endpoint,
	}
}
