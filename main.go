package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/amaraliou/spotistats/api"
	"github.com/amaraliou/spotistats/handlers"
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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", handlers.HandleHomepage)
	http.HandleFunc("/login", api.HandleLoginRequest)
	http.HandleFunc("/callback", api.CallbackHandler)
	fmt.Println(http.ListenAndServe(":8000", nil))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}
	t.ExecuteTemplate(w, "base", nil)
}
