package api

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	BaseAccountURL = "accounts.spotify.com/"
)

type Token struct {
	AccessToken    string        `json:"access_token"`
	TokenType      string        `json:"token_type"`
	ExpiresIn      time.Duration `json:"expires_in"`
	ExpirationDate time.Time     `json:"expiration_date"`
	ClientID       string        `json:"client_id"`
	ClientSecret   string        `json:"client_secret"`
	RefreshToken   string        `json:"refresh_token"`
	Scope          string        `json:"scope"`
}

var token Token

func serverManager(srv *http.Server, keepAlive chan bool) {
	for {
		select {
		case <-keepAlive:
			ctx := context.Background()

			err := srv.Shutdown(ctx)

			if err != nil {
				fmt.Printf("Failed to shutdown temporary http server, this should have no effect on your ability to complete this process.  It's running on port 15298.")
			}

			return
		default:
		}
	}
}

func GetAuthorizationURL(id string) string {
	v := url.Values{}
	v.Set("client_id", id)
	v.Set("response_type", "code")
	v.Set("redirect_uri", "http://localhost:8000/callback")
	v.Set("scope", "playlist-read-private user-top-read user-library-read user-library-modify user-read-currently-playing user-read-recently-played user-modify-playback-state user-read-playback-state user-follow-read playlist-read-collaborative")

	r := buildReq("GET", BaseAccountURL+"authorize", v, nil)
	return r.URL.String()
}

func CodeAuthorize(id, secret, code string) Token {
	var t Token
	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", "http://localhost:8000/callback")

	r := buildReq("POST", BaseAccountURL+"api/token", v, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.SetBasicAuth(id, secret)

	err := makeReq(r, &t)

	if err != nil {
		log.Fatal(err)
	}

	t.ExpirationDate = time.Now().Add((t.ExpiresIn - 30) * time.Second)
	t.ClientID = id
	t.ClientSecret = secret

	return t
}

func GetCode(id string) (c string) {
	m := GetAuthorizationURL(id)
	fmt.Printf("\nNavigate to the following URL to Authorize Spotistats:\n%s\n", m)
	keepAlive := make(chan bool)
	server := &http.Server{Addr: ":8000"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		keepAlive <- true
		fmt.Fprintf(w, "<h1>Almost done!</h1><p>Spotistats has been approved, just copy the following code back to the CLI:\n <span style=\"color: #FF0000\">%s</span></p>", code)
	})

	go serverManager(server, keepAlive)
	server.ListenAndServe()

	fmt.Print("\nEnter Code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c = strings.TrimSpace(scanner.Text())

	return c
}

func Authenticate(id, secret string) {
	code := GetCode(id)
	token = CodeAuthorize(id, secret, code)
	fmt.Println("\nSuccessful Authentication")
}
