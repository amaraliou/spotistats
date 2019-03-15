package main

import (
	"fmt"
	"log"

	"github.com/amaraliou/spotistats/api"
)

func main() {
	api.Authenticate("", "")
	track, err := api.GetTrack("5jwnLyVoVAxANd9cVYjHws")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Track: %s\n", track.Name)
	fmt.Printf("Artist: %s\n", track.Artists[0].Name)
	fmt.Printf("Album: %s\n", track.Album.Name)
	fmt.Printf("Duration: %s\n", api.MillisecsToSongTime(track.Duration))
}
