package api

import "net/http"

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
