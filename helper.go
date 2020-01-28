package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	URL     = "https://www.flickr.com/services/rest"
	client  = &http.Client{}
	API_KEY = os.Getenv("FLICKR_API_KEY")
)

// func main() {
// GetUserIDByUsername("Acoustics & Media Society, IIIT Allahabad")
// GetAlbumsFromUserID("129074767@N06")
// GetPhotosFromAlbum("129074767@N06", "72157712783898977")
// GetPhotoSizes("49424836736")
// DownloadPhoto("https://live.staticflickr.com/65535/49424836736_94cdb4d91d_b.jpg", "./assets/demo.jpg")
// }

func GetUserIDByUsername(username string) (string, error) {

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", errors.New("API Request Failed")
	}

	q := req.URL.Query()
	q.Add("method", "flickr.people.findByUsername")
	q.Add("api_key", API_KEY)
	q.Add("username", username)
	q.Add("format", "json")
	q.Add("nojsoncallback", "1")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New("API Request Failed")
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New("The API returned unexpected values")
	}

	var user UserStruct
	json.Unmarshal(data, &user)

	if user.User.ID == "" {
		return "", errors.New("Can't find a user with username " + username)
	}

	return user.User.ID, nil
}

func GetAlbumsFromUserID(userID string) ([]string, error) {

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return []string{}, errors.New("Can't make a request to the API")
	}

	q := req.URL.Query()
	q.Add("method", "flickr.photosets.GetList")
	q.Add("api_key", API_KEY)
	q.Add("user_id", userID)
	q.Add("format", "json")
	q.Add("nojsoncallback", "1")

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		return []string{}, errors.New("Something wrong with API")
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []string{}, errors.New("Can't read the response body")
	}

	var album Album
	json.Unmarshal(data, &album)

	var photos []string
	for _, a := range album.Photosets.Photoset {
		photos = append(photos, a.Title.Content)
	}

	return photos, nil
}

func GetPhotosFromAlbum(userID string, photosetID string) (map[string]string, error) {

	req, _ := http.NewRequest("GET", URL, nil)

	q := req.URL.Query()
	q.Add("method", "flickr.photosets.GetPhotos")
	q.Add("api_key", API_KEY)
	q.Add("user_id", userID)
	q.Add("photoset_id", photosetID)
	q.Add("format", "json")
	q.Add("nojsoncallback", "1")

	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var photos AlbumPhotos
	json.Unmarshal(data, &photos)

	list_of_photos := make(map[string]string)
	for _, photo := range photos.Photoset.Photo {
		list_of_photos[photo.ID] = photo.Title + ".jpg"
	}

	return list_of_photos, nil
}

func GetPhotoSizes(photoID string) (map[string]string, error) {

	req, _ := http.NewRequest("GET", URL, nil)

	q := req.URL.Query()
	q.Add("method", "flickr.photos.GetSizes")
	q.Add("api_key", API_KEY)
	q.Add("photo_id", photoID)
	q.Add("format", "json")
	q.Add("nojsoncallback", "1")

	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var sizes PhotoSizes
	json.Unmarshal(data, &sizes)

	links := make(map[string]string)
	for _, photo := range sizes.Sizes.Size {
		links[photo.Source] = photo.Label
	}

	return links, nil
}

func DownloadPhoto(url string, filename string) error {

	resp, _ := http.Get(url)

	file, _ := os.Create(filename)
	size, _ := io.Copy(file, resp.Body)

	fmt.Println(size)

	return nil
}
