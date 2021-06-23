package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"photo-album-assignment/controller"
	"photo-album-assignment/external_client"
	"photo-album-assignment/repository"
)

func main(){
	fmt.Println("Starting Up Application")
	// to run always
	repository.InitDatabase()

	if toScrapData() {
		PopulateData()
	}else{
		ServeRequest(controller.GetRoutes())
	}
}

func toScrapData() bool {
	return os.Getenv("MODE") == "scrapper"
}

func ServeRequest(routes *mux.Router) {
	log.Println("Starting Up Server on port 8080")
	err := http.ListenAndServe(
		"localhost:8080",
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.MaxAge(600),
		)(routes))
	if err != nil {
		log.Fatalln("Error starting server , error : ", err)
	}
	os.Exit(1)
}



func PopulateData() {
	log.Println("Starting Data Scrapper after a second")
	time.Sleep(time.Second)
	repository.ResetTables()
	albums := external_client.FetchAlbums()
	fmt.Println(len(albums))
	for _, album := range albums {
		albumPhotos := external_client.FetchPhotos(album.ID)
		log.Println("Found", len(albumPhotos), "Photos for:", album.Title)
		repository.InsertData(album, albumPhotos)
		time.Sleep(100 * time.Millisecond)
	}
	log.Println("Data Populated")
}

