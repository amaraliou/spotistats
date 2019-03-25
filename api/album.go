package api

import (
	"strconv"
	"strings"
	"time"
)

type BaseAlbum struct {
	AlbumType    string            `json:"album"`
	Artists      []BaseArtist      `json:"artists"`
	Markets      []string          `json:"available_markets"`
	ExternalURLS map[string]string `json:"external_urls"`
	URI          string            `json:"uri"`
	Href         string            `json:"href"`
	Type         string            `json:"type"`
	ID           string            `json:"id"`
	Images       []Image           `json:"images"`
	Name         string            `json:"name"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type FullAlbum struct {
	BaseAlbum
	Tracks               []BaseTrack `json:"tracks"`
	ReleaseDate          string      `json:"release_date"`
	ReleaseDatePrecision string      `json:"release_date_precision"`
	Popularity           int         `json:"popularity"`
	Label                string      `json:"label"`
	Genres               []string    `json:"genres"`
	Copyrights           []Copyright `json:"copyrights"`
}

type FullAlbums struct {
	Albums []FullAlbum `json:"albums"`
}

type SavedAlbum struct {
	AddedAt *time.Time `json:"added_at"`
	Album   BaseAlbum  `json:"album"`
}

type AlbumList struct {
	Href     string      `json:"href"`
	Items    []BaseAlbum `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

type SavedAlbumList struct {
	Href     string       `json:"href"`
	Items    []SavedAlbum `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

func GetAlbum(albumID string) (album FullAlbum, err error) {

	r := buildReq("GET", BaseURL+"albums/"+albumID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &album)
	return album, err
}

func GetMultipleAlbums(albumIDs ...string) (albums FullAlbums, err error) {
	IDs := strings.Join(albumIDs, ",")
	r := buildReq("GET", BaseURL+"albums/?ids="+IDs, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &albums)
	return albums, err
}

//To add offset and limit (optionals)
func GetAlbumTracks(albumID string, limit, offset int) (tracksPage TrackList, err error) {

	r := buildReq("GET", BaseURL+"albums/"+albumID+"/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	q := r.URL.Query()

	if limit != 20 {
		q.Add("limit", strconv.Itoa(limit))
	}

	if offset != 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	r.URL.RawQuery = q.Encode()

	err = makeReq(r, &tracksPage)
	return tracksPage, err
}
