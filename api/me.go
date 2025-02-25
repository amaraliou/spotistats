package api

import (
	"net/http"
	"net/url"
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

type AllTopTracks struct {
	Short  TopTracks `json:"short_term"`
	Medium TopTracks `json:"medium_term"`
	Long   TopTracks `json:"long_term"`
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

type AllTopArtists struct {
	Short  TopArtists `json:"short_term"`
	Medium TopArtists `json:"medium_term"`
	Long   TopArtists `json:"long_term"`
}

//To edit datetime object for Birthdate
type UserInfo struct {
	DisplayName  string            `json:"display_name"`
	Birthdate    string            `json:"birthdate"`
	Country      string            `json:"country"`
	Email        string            `json:"email"`
	ExternalURLS map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Followers    Followers         `json:"followers"`
	Images       []Image           `json:"images"`
	Product      string            `json:"product"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type Device struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Active     bool   `json:"is_active"`
	Private    bool   `json:"is_private_session"`
	Restricted bool   `json:"is_restricted"`
	Type       string `json:"type"`
	Volume     int    `json:"volume_percent"`
}

type Context struct {
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
	ExternalURLS map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
}

type CurrentPlayback struct {
	Device    Device    `json:"device"`
	Repeat    string    `json:"repeat_state"`
	Shuffle   bool      `json:"shuffle_state"`
	Context   Context   `json:"context"`
	Timestamp int       `json:"timestamp"`
	Progress  int       `json:"progress_ms"`
	Playing   bool      `json:"is_playing"`
	Item      FullTrack `json:"item"`
	Type      string    `json:"type"`
}

func GetSavedAlbums(limit, offset int, token string) (savedAlbums SavedAlbumList, err error) {

	q := url.Values{}

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r := buildReq("GET", BaseURL+"me/albums", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &savedAlbums)
	return savedAlbums, err
}

func GetNextSavedAlbums(url string, token string) (savedAlbums SavedAlbumList, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &savedAlbums)
	return savedAlbums, err
}

func GetSavedTracks(limit, offset int, token string) (savedTracks SavedTrackList, err error) {

	q := url.Values{}

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r := buildReq("GET", BaseURL+"me/tracks", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &savedTracks)
	return savedTracks, err
}

func GetNextSavedTracks(url string, token string) (savedTracks SavedTrackList, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &savedTracks)
	return savedTracks, err
}

func GetTopTracks(timeRange string, limit, offset int, token string) (topTracks TopTracks, err error) {

	q := url.Values{}
	q.Add("time_range", timeRange)

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r := buildReq("GET", BaseURL+"me/top/tracks", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &topTracks)
	return topTracks, err
}

func GetNextTopTracks(url string, token string) (topTracks TopTracks, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &topTracks)
	return topTracks, err
}

func GetAllTopTracks(token string) (allTopTracks AllTopTracks, err error) {

	allTopTracks.Short, err = GetTopTracks("short_term", 50, 0, token)
	allTopTracks.Medium, err = GetTopTracks("medium_term", 50, 0, token)
	allTopTracks.Long, err = GetTopTracks("long_term", 50, 0, token)

	return allTopTracks, err
}

func GetTopArtists(timeRange string, limit, offset int, token string) (topArtists TopArtists, err error) {

	q := url.Values{}
	q.Add("time_range", timeRange)

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r := buildReq("GET", BaseURL+"me/top/artists", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	err = makeReq(r, &topArtists)
	return topArtists, err
}

func GetNextTopArtists(url string, token string) (topArtists TopArtists, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &topArtists)
	return topArtists, err
}

func GetAllTopArtists(token string) (allTopArtists AllTopArtists, err error) {

	allTopArtists.Short, err = GetTopArtists("short_term", 50, 0, token)
	allTopArtists.Medium, err = GetTopArtists("medium_term", 50, 0, token)
	allTopArtists.Long, err = GetTopArtists("long_term", 50, 0, token)

	return allTopArtists, err
}

func GetMyInfo(token string) (myInfo UserInfo, err error) {

	r := buildReq("GET", BaseURL+"me", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &myInfo)
	return myInfo, err
}

func GetCurrentPlayback(token string) (playback CurrentPlayback, err error) {

	r := buildReq("GET", BaseURL+"me/player", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &playback)
	return playback, err
}
