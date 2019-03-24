package api

import "net/http"

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
