package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"photo-album-assignment/repository"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	c := ApiController{repository.DBScope}
	r.HandleFunc("/search", c.SearchData).Methods(http.MethodGet)
	return r
}
