package main

import (
	"fmt"
	"log"

	"github.com/amaraliou/spotistats/api"
)

func main() {
	api.Authenticate("", "")
	topTracks, err := api.GetTopTracks()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTOP TRACKS:\n")
	for _, track := range topTracks.Items {
		fmt.Printf("\n%s", track.Name)
	}
	fmt.Printf("\n")
}
