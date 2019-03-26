package api

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	spotifyOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	spotifyOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     "",
		ClientSecret: "",
		Scopes: []string{
			"playlist-read-private",
			"user-top-read",
			"user-library-read",
			"user-library-modify",
			"user-read-currently-playing",
			"user-read-recently-played",
			"user-modify-playback-state",
			"user-read-playback-state",
			"user-follow-read",
			"playlist-read-collaborative",
		},
		Endpoint: spotify.Endpoint,
	}
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/login">Spotify Log In</a>
</body>
</html>`

	fmt.Fprintf(w, htmlIndex)
}

func HandleSpotifyLogin(w http.ResponseWriter, r *http.Request) {
	url := spotifyOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleSpotifyCallback(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", token.AccessToken)
}

func getToken(state string, code string) (*oauth2.Token, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := spotifyOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	return token, err
}
