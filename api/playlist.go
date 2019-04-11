package api

import "time"

type PublicUser struct {
	DisplayName  string            `json:"display_name"`
	ExternalURLS map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Followers    Followers         `json:"followers"`
	Images       []Image           `json:"images"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type PlaylistTracksLink struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type PlaylistTrack struct {
	AddedAt time.Time  `json:"added_at"`
	AddedBy PublicUser `json:"added_by"`
	IsLocal bool       `json:"is_local"`
	Track   FullTrack  `json:"track"`
}

type BasePlaylist struct {
	Name          string             `json:"name"`
	URI           string             `json:"uri"`
	Href          string             `json:"href"`
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	ExternalURLS  map[string]string  `json:"external_urls"`
	Images        []Image            `json:"images"`
	Collaborative bool               `json:"collaborative"`
	Owner         PublicUser         `json:"owner"`
	Public        bool               `json:"public"`
	SnapshotID    string             `json:"snapshot_id"`
	Tracks        PlaylistTracksLink `json:"tracks"`
}

type FullPlaylist struct {
	Name          string            `json:"name"`
	URI           string            `json:"uri"`
	Href          string            `json:"href"`
	Type          string            `json:"type"`
	ID            string            `json:"id"`
	ExternalURLS  map[string]string `json:"external_urls"`
	Images        []Image           `json:"images"`
	Collaborative bool              `json:"collaborative"`
	Owner         PublicUser        `json:"owner"`
	Followers     Followers         `json:"followers"`
	Public        bool              `json:"public"`
	SnapshotID    string            `json:"snapshot_id"`
	Tracks        PlaylistTrackList `json:"tracks"`
	Description   string            `json:"description"`
}

type PlaylistList struct {
	Href     string         `json:"href"`
	Items    []BasePlaylist `json:"items"`
	Limit    int            `json:"limit"`
	Next     string         `json:"next"`
	Offset   int            `json:"offset"`
	Previous string         `json:"previous"`
	Total    int            `json:"total"`
}

type PlaylistTrackList struct {
	Href     string          `json:"href"`
	Items    []PlaylistTrack `json:"items"`
	Limit    int             `json:"limit"`
	Next     string          `json:"next"`
	Offset   int             `json:"offset"`
	Previous string          `json:"previous"`
	Total    int             `json:"total"`
}
