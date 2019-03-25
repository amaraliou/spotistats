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

type Features struct {
	Duration         int     `json:"duration_ms"`
	Key              int     `json:"key"`
	Mode             int     `json:"mode"`
	Signature        int     `json:"time_signature"`
	Acousticness     float32 `json:"acousticness"`
	Danceability     float32 `json:"danceability"`
	Energy           float32 `json:"energy"`
	Instrumentalness float32 `json:"instrumentalness"`
	Liveness         float32 `json:"liveness"`
	Loudness         float32 `json:"loudness"`
	Speechiness      float32 `json:"speechiness"`
	Valence          float32 `json:"valence"`
	Tempo            float32 `json:"tempo"`
	Href             string  `json:"track_href"`
	ID               string  `json:"id"`
	URI              string  `json:"uri"`
	Type             string  `json:"type"`
	AnalysisURL      string  `json:"analysis_url"`
}

type CurrentTrack struct {
	Context   Context   `json:"context"`
	Timestamp time.Time `json:"timestamp"`
	Progress  int       `json:"progress_ms"`
	Playing   bool      `json:"is_playing"`
	Item      FullTrack `json:"item"`
	Type      string    `json:"current_playing_type"`
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

//BETA
func GetCurrentTrack(market string) (currentTrack CurrentTrack, err error) {

	r := buildReq("GET", BaseURL+"me/currently-playing", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	q := r.URL.Query()
	q.Add("market", market)
	r.URL.RawQuery = q.Encode()

	err = makeReq(r, &currentTrack)
	return currentTrack, err
}

func GetAudioFeatures(trackID string) (audioFeatures Features, err error) {

	r := buildReq("GET", BaseURL+"audio-features/"+trackID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &audioFeatures)
	return audioFeatures, err
}
