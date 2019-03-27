package api

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	OAuthConfig *oauth2.Config

	SessionStore sessions.Store
)

const (
	ScopeImageUpload               = "ugc-image-upload"
	ScopePlaylistReadPrivate       = "playlist-read-private"
	ScopePlaylistModifyPublic      = "playlist-modify-public"
	ScopePlaylistModifyPrivate     = "playlist-modify-private"
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	ScopeUserFollowModify          = "user-follow-modify"
	ScopeUserFollowRead            = "user-follow-read"
	ScopeUserLibraryModify         = "user-library-modify"
	ScopeUserLibraryRead           = "user-library-read"
	ScopeUserReadPrivate           = "user-read-private"
	ScopeUserReadEmail             = "user-read-email"
	ScopeUserReadBirthdate         = "user-read-birthdate"
	ScopeUserReadCurrentlyPlaying  = "user-read-currently-playing"
	ScopeUserReadPlaybackState     = "user-read-playback-state"
	ScopeUserModifyPlaybackState   = "user-modify-playback-state"
	ScopeUserReadRecentlyPlayed    = "user-read-recently-played"
	ScopeUserTopRead               = "user-top-read"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	OAuthConfig = configureOAuthClient(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	cookieStore := sessions.NewCookieStore([]byte(""))
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
