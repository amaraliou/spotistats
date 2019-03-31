package main

import (
	"fmt"
	"html/template"
	"log"
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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", HandleHome)
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

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("templates/base-logged-in.html", "templates/navbar.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}

	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)

	topArtists, err := api.GetTopArtists("long_term", 10, 0, tok)
	topTracks, err := api.GetTopTracks("long_term", 10, 0, tok)
	data := map[string]interface{}{
		"TopArtists": topArtists,
		"TopTracks":  topTracks,
	}
	t.ExecuteTemplate(writer, "base-logged-in", data)
}
