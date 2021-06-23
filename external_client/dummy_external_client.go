package external_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"photo-album-assignment/models"
)

const(
	albumsEndPoint = "https://jsonplaceholder.typicode.com/albums"
	photosEndPoint = "https://jsonplaceholder.typicode.com/photos?albumId=%s"
)

func FetchAlbums() []models.Album {
	resp, err := http.Get(albumsEndPoint)
	if err != nil {
		log.Println("Failed to fetch albums | error : ", err)
		return []models.Album{}
	}
	defer resp.Body.Close()
	var albums []models.Album
	err = json.NewDecoder(resp.Body).Decode(&albums)
	if err != nil {
		log.Println("Failed to parse albums response | error : ", err)
		return []models.Album{}
	}
	return albums
}


func FetchPhotos(albumId int) []models.Photo {
	resp, err := http.Get(fmt.Sprintf(photosEndPoint, strconv.Itoa(albumId)))
	if err != nil {
		log.Println("Failed to fetch photos | error : ", err)
		return []models.Photo{}
	}
	defer resp.Body.Close()
	var photos []models.Photo
	err = json.NewDecoder(resp.Body).Decode(&photos)
	if err != nil {
		log.Println("Failed to parse photos response | error : ", err)
		return []models.Photo{}
	}
	return photos
}


