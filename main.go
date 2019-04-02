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
	http.HandleFunc("/top/tracks", HandleTopTracks)
	http.HandleFunc("/top/artists", HandleTopArtists)
	http.HandleFunc("/recent", HandleRecentTracks)
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
	t, err := template.ParseFiles("templates/homepage.html", "templates/navbar.html", "templates/floatbutton.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}

	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)

	topArtists, err := api.GetTopArtists("long_term", 10, 0, tok)
	if err != nil {
		log.Fatal(err)
	}

	topTracks, err := api.GetTopTracks("long_term", 10, 0, tok)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := api.GetMyInfo(tok)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"TopArtists": topArtists,
		"TopTracks":  topTracks,
		"User":       userInfo,
	}
	t.ExecuteTemplate(writer, "homepage", data)
}

func HandleTopTracks(writer http.ResponseWriter, request *http.Request) {
	t, err := template.New("templates/toptracks.html").Funcs(
		template.FuncMap{
			"add": func(a int) int {
				return a + 1
			},
		},
	).ParseFiles("templates/toptracks.html", "templates/navbar.html", "templates/floatbutton.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}

	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)
	allTopTracks, err := api.GetAllTopTracks(tok)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"longTerm":   allTopTracks.Long,
		"mediumTerm": allTopTracks.Medium,
		"shortTerm":  allTopTracks.Short,
	}
	t.ExecuteTemplate(writer, "toptracks", data)
}

func HandleTopArtists(writer http.ResponseWriter, request *http.Request) {
	t, err := template.New("templates/topartists.html").Funcs(
		template.FuncMap{
			"add": func(a int) int {
				return a + 1
			},
		},
	).ParseFiles("templates/topartists.html", "templates/navbar.html", "templates/floatbutton.html")

	if err != nil {
		log.Fatal(err)
	}

	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)
	allTopArtists, err := api.GetAllTopArtists(tok)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"longTerm":   allTopArtists.Long,
		"mediumTerm": allTopArtists.Medium,
		"shortTerm":  allTopArtists.Short,
	}

	t.ExecuteTemplate(writer, "topartists", data)
}

func HandleRecentTracks(writer http.ResponseWriter, request *http.Request) {
	t, err := template.New("templates/recent.html").Funcs(
		template.FuncMap{
			"add": func(a int) int {
				return a + 1
			},
		},
	).ParseFiles("templates/recent.html", "templates/navbar.html", "templates/floatbutton.html")

	if err != nil {
		log.Fatal(err)
	}

	currentSession, err := api.SessionStore.Get(request, "spotistats")
	if err != nil {
		log.Fatal(err)
	}

	tok := currentSession.Values["oauth_token"].(string)
	recentFullTracks, err := api.GetRecentFullTracks(50, tok)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"tracks": recentFullTracks,
	}

	t.ExecuteTemplate(writer, "recent", data)
}
