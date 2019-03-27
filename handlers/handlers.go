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
	topTracks, err := api.GetTopTracks("long_term", 10, 0, tok)
	if err != nil {
		log.Fatal(err)
	}

	topArtists, err := api.GetTopArtists("long_term", 10, 0, tok)
	if err != nil {
		log.Fatal(err)
	}

	currentlyPlaying, err := api.GetCurrentTrack("GB", tok)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(writer, "<h1>CURRENTLY PLAYING</h1>")
	fmt.Fprintf(writer, "<h3>%s</h3>", currentlyPlaying.Item.Name)

	fmt.Fprintf(writer, "<h1>TOP TRACKS</h1>")
	for number, track := range topTracks.Items {
		fmt.Fprintf(writer, "<h3>%d %s</h3>",
			number+1,
			track.Name)
	}

	fmt.Fprintf(writer, "<h1>TOP ARTISTS</h1>")
	for number, artist := range topArtists.Items {
		fmt.Fprintf(writer, "<h3>%d %s</h3>",
			number+1,
			artist.Name)
	}
}

func HandleTopTracksPage(writer http.ResponseWriter, request *http.Request) {

}

func HandleTopArtistsPage(writer http.ResponseWriter, request *http.Request) {

}

func HandleRecentlyPlayedPage(writer http.ResponseWriter, request *http.Request) {

}
