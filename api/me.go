package api

import (
	"net/http"
	"strconv"
)

type TopTracks struct {
	Href     string      `json:"href"`
	Items    []FullTrack `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

type TopArtists struct {
	Href     string       `json:"href"`
	Items    []FullArtist `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

//To add offset and limit (optionals)
func GetSavedAlbums() (savedAlbums SavedAlbumList, err error) {

	r := buildReq("GET", BaseURL+"me/albums", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &savedAlbums)
	return savedAlbums, err
}

func GetNextSavedAlbums(url string) (savedAlbums SavedAlbumList, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &savedAlbums)
	return savedAlbums, err
}

//To add offset and limit (optionals)
func GetSavedTracks() (savedTracks SavedTrackList, err error) {

	r := buildReq("GET", BaseURL+"me/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &savedTracks)
	return savedTracks, err
}

func GetNextSavedTracks(url string) (savedTracks SavedTrackList, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &savedTracks)
	return savedTracks, err
}

func GetTopTracks(timeRange string, limit, offset int) (topTracks TopTracks, err error) {

	r := buildReq("GET", BaseURL+"me/top/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)
	q := r.URL.Query()
	q.Add("time_range", timeRange)

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r.URL.RawQuery = q.Encode()

	err = makeReq(r, &topTracks)
	return topTracks, err
}

func GetNextTopTracks(url string) (topTracks TopTracks, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &topTracks)
	return topTracks, err
}

func GetTopArtists() (topArtists TopArtists, err error) {

	r := buildReq("GET", BaseURL+"me/top/artists", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &topArtists)
	return topArtists, err
}

func GetNextTopArtists(url string) (topArtists TopArtists, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &topArtists)
	return topArtists, err
}
