package main

import (
	"fmt"
	"log"

	"github.com/amaraliou/spotistats/api"
)

func main() {
	api.Authenticate("", "")
	topTracks, err := api.GetTopTracks("medium_term")

	if err != nil {
		log.Fatal(err)
	}

	for _, track := range topTracks.Items {
		fmt.Printf("\n%s", track.Name)
	}
}
