package handlers

import (
	"log"
	"net/http"

	"github.com/amaraliou/spotistats/api"
)

func HandleHomepage(writer http.ResponseWriter, request *http.Request) {
	oauthFlowSession, err := api.SessionStore.Get(request, request.FormValue("state"))
	if err != nil {
		log.Fatal(err)
	}
}
