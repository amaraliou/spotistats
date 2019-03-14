package api

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

type TrackList struct {
	Href     string      `json:"href"`
	Items    []FullTrack `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

func GetTrack(trackID string) (track FullTrack, err error) {
	r := buildReq("GET", BaseURL+"tracks/"+trackID, nil, nil)
	r.Header.Add("Authorization", "Bearer "+token.AccessToken)

	err = makeReq(r, &track)

	return track, err
}
