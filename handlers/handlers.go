package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amaraliou/spotistats/api"
)

func HandleHomepage(writer http.ResponseWriter, request *http.Request) {
	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)
	fmt.Print(tok)
	topTracks, err := api.GetTopTracks("short_term", 50, 0, tok)
	if err != nil {
		log.Fatal(err)
	}

	for number, track := range topTracks.Items {
		fmt.Fprintf(writer, "<h1>%d %s</h1>",
			number+1,
			track.Name)
	}
}
