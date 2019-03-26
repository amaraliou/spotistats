package api

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/oauth2"
)

const (
	defaultSessionID     = "default"
	oauthTokenSessionKey = "oauth_token"
	oauthFlowRedirectKey = "redirect"
)

func init() {
	gob.Register(&oauth2.Token{})
}

func HandleLoginRequest(writer http.ResponseWriter, request *http.Request) {
	sessionID := uuid.Must(uuid.NewV4()).String()

	oauthFlowSession, err := SessionStore.New(request, sessionID)
	if err != nil {
		log.Fatal(err)
	}

	oauthFlowSession.Options.MaxAge = 3600

	oauthFlowSession.Values[oauthFlowRedirectKey] = "/home"

	if err := oauthFlowSession.Save(request, writer); err != nil {
		log.Fatal(err)
	}

	authCodeURL := OAuthConfig.AuthCodeURL(sessionID, oauth2.ApprovalForce, oauth2.AccessTypeOnline)
	http.Redirect(writer, request, authCodeURL, http.StatusFound)
}

func HandleLogoutRequest(writer http.ResponseWriter, request *http.Request) {
	session, err := SessionStore.New(request, defaultSessionID)
	if err != nil {
		log.Fatal(err)
	}
	session.Options.MaxAge = -1
	if err := session.Save(request, writer); err != nil {
		log.Fatal(err)
	}
	redirectURL := request.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/"
	}
	http.Redirect(writer, request, redirectURL, http.StatusFound)
}

func CallbackHandler(writer http.ResponseWriter, request *http.Request) {
	oauthFlowSession, err := SessionStore.Get(request, request.FormValue("state"))
	if err != nil {
		log.Fatal(err)
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)

	if !ok {
		log.Fatal(err)
	}

	code := request.FormValue("code")
	tok, err := OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	session, err := SessionStore.New(request, defaultSessionID)
	if err != nil {
		log.Fatal(err)
	}

	session.Values[oauthTokenSessionKey] = tok
	if err := session.Save(request, writer); err != nil {
		log.Fatal(err)
	}

	http.Redirect(writer, request, redirectURL, http.StatusFound)
}
