package api

import "time"

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
