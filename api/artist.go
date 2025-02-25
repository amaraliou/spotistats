package api

import (
	"net/http"
	"strings"
)

type Followers struct {
	Total int `json:"total"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type BaseArtist struct {
	Name         string            `json:"name"`
	ID           string            `json:"id"`
	URI          string            `json:"uri"`
	Href         string            `json:"href"`
	ExternalURLS map[string]string `json:"external_urls"`
	Type         string            `json:"type"`
}

type FullArtist struct {
	Name         string            `json:"name"`
	ID           string            `json:"id"`
	URI          string            `json:"uri"`
	Href         string            `json:"href"`
	ExternalURLS map[string]string `json:"external_urls"`
	Type         string            `json:"type"`
	Popularity   int               `json:"popularity"`
	Images       []Image           `json:"images"`
	Genres       []string          `json:"genres"`
	Followers    *Followers        `json:"followers"`
}

type FullArtists struct {
	Artists []FullArtist `json:"artists"`
}

type ArtistList struct {
	Href     string       `json:"href"`
	Items    []FullArtist `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

func GetArtistAlbums(artistID string, token string) (albums AlbumList, err error) {

	r := buildReq("GET", BaseURL+"artists/"+artistID+"/albums", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &albums)

	return albums, err
}

func GetNextArtistAlbums(url string, token string) (albums AlbumList, err error) {

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &albums)

	return albums, err
}

func GetArtist(artistID string, token string) (artist FullArtist, err error) {

	r := buildReq("GET", BaseURL+"artists/"+artistID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &artist)
	return artist, err
}

func GetMultipleArtists(token string, artistIDs ...string) (artists FullArtists, err error) {
	IDs := strings.Join(artistIDs, ",")
	r := buildReq("GET", BaseURL+"artists/?ids="+IDs, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &artists)

	return artists, err
}

func GetArtistTopTracks(artistID string, token string) (topTracks FullTracks, err error) {

	r := buildReq("GET", BaseURL+"artists/"+artistID+"/top-tracks?country=GB", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &topTracks)
	return topTracks, err
}

func GetArtistRelatedArtists(artistID string, token string) (relatedArtists FullArtists, err error) {

	r := buildReq("GET", BaseURL+"artists/"+artistID+"/related-artists", nil, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	err = makeReq(r, &relatedArtists)
	return relatedArtists, err
}
