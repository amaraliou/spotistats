package api

import (
	"strings"
	"time"
)

type BaseTrack struct {
	Artists      []BaseArtist      `json:"artists"`
	Markets      []string          `json:"available_markets"`
	DiscNumber   int               `json:"disc_number"`
	Duration     int               `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalURLS map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type FullTrack struct {
	BaseTrack
	Album      *BaseAlbum `json:"album"`
	Popularity int        `json:"popularity"`
}

type SavedTrack struct {
	AddedAt *time.Time `json:"added_at"`
	Track   BaseTrack  `json:"track"`
}

type SavedTrackList struct {
	Href     string       `json:"href"`
	Items    []SavedTrack `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

type TrackList struct {
	Href     string      `json:"href"`
	Items    []BaseTrack `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

type FullTracks struct {
	Tracks []FullTrack `json:"tracks"`
}

func GetTrack(trackID string) (track FullTrack, err error) {
	r := buildReq("GET", BaseURL+"tracks/"+trackID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &track)

	return track, err
}

func GetMultipleTracks(trackIDs ...string) (tracks FullTracks, err error) {
	IDs := strings.Join(trackIDs, ",")
	r := buildReq("GET", BaseURL+"tracks/?ids="+IDs, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &tracks)

	return tracks, err
}

func getCurrentTrack(market string) (currentTrack CurrentTrack, err error) {

	r := buildReq("GET", BaseURL+"me/currently-playing", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	q := r.URL.Query()
	q.Add("market", market)
	r.URL.RawQuery = q.Encode()

	err = makeReq(r, &currentTrack)
	return currentTrack, err
}
