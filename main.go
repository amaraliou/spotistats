package main

import (
	"fmt"
	"net/http"

	"github.com/amaraliou/spotistats/api"
)

func main() {
	//api.Authenticate("", "")
	//topTracks, err := api.GetTopTracks("short_term", 50, 0)

	//if err != nil {
	//log.Fatal(err)
	//}

	//fmt.Printf("\nTOP TRACKS:\n")
	//for number, track := range topTracks.Items {
	//fmt.Printf("\n%d - %s", number+1, track.Name)
	//}
	//fmt.Printf("\n")

	http.HandleFunc("/", api.HandleMain)
	http.HandleFunc("/login", api.HandleSpotifyLogin)
	http.HandleFunc("/callback", api.HandleSpotifyCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))

}
