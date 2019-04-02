package api

import (
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
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
	Album      BaseAlbum `json:"album"`
	Popularity int       `json:"popularity"`
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

//To change timestamp type
type CurrentTrack struct {
	Context   Context   `json:"context"`
	Timestamp int       `json:"timestamp"`
	Progress  int       `json:"progress_ms"`
	Playing   bool      `json:"is_playing"`
	Item      FullTrack `json:"item"`
	Type      string    `json:"current_playing_type"`
}

type RecentTrack struct {
	Track    BaseTrack `json:"track"`
	PlayedAt time.Time `json:"played_at"`
	Context  Context   `json:"context"`
}

type Cursor struct {
	Before time.Time `json:"before"`
	After  time.Time `json:"after"`
}

type RecentTrackList struct {
	Href   string        `json:"href"`
	Items  []RecentTrack `json:"items"`
	Limit  int           `json:"limit"`
	Next   string        `json:"next"`
	Cursor Cursor        `json:"cursor"`
	Total  int           `json:"total"`
}

func GetTrack(trackID string, token string) (track FullTrack, err error) {
	r := buildReq("GET", BaseURL+"tracks/"+trackID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &track)

	return track, err
}

func GetMultipleTracks(token string, trackIDs ...string) (tracks FullTracks, err error) {
	IDs := strings.Join(trackIDs, ",")
	r := buildReq("GET", BaseURL+"tracks/?ids="+IDs, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &tracks)

	return tracks, err
}

func GetCurrentTrack(market string, token string) (currentTrack CurrentTrack, err error) {

	q := url.Values{}
	q.Add("market", market)

	r := buildReq("GET", BaseURL+"me/player/currently-playing", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &currentTrack)
	return currentTrack, err
}

func GetAudioFeatures(trackID string, token *oauth2.Token) (audioFeatures Features, err error) {

	r := buildReq("GET", BaseURL+"audio-features/"+trackID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &audioFeatures)
	return audioFeatures, err
}

func GetRecentTracks(limit int, token string) (recentTracks RecentTrackList, err error) {

	q := url.Values{}

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	r := buildReq("GET", BaseURL+"me/player/recently-played", q, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &recentTracks)

	return recentTracks, err
}

func GetRecentFullTracks(limit int, token string) (recentFullTracks []FullTrack, err error) {

	recentTracks, err := GetRecentTracks(limit, token)
	if err != nil {
		log.Fatal(err)
	}

	for _, recentTrackItem := range recentTracks.Items {
		id := recentTrackItem.Track.ID
		fullTrack, err := GetTrack(id, token)
		if err != nil {
			log.Fatal(err)
		}
		recentFullTracks = append(recentFullTracks, fullTrack)
	}

	return recentFullTracks, err
}
