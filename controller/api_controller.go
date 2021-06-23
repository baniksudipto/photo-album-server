package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"gorm.io/gorm"

	"photo-album-assignment/models"
)

type ApiController struct {
	scope *gorm.DB
}

func (c ApiController) SearchData(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		RenderResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	queryParams := new(models.QueryParams)
	if err := schema.NewDecoder().Decode(queryParams, r.Form); err != nil {
		RenderResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	log.Printf("Got Request For %+v", queryParams)
	switch queryParams.Type {
	case "album":{
		c.FindAlbum(w, queryParams)
	}
	case "photo":
		c.FindPhoto(w, queryParams)
	default:
		RenderResponse(w, http.StatusUnsupportedMediaType, map[string]string{"error": queryParams.Type + " unknown"})
	}

}

func (c ApiController) FindPhoto(w http.ResponseWriter, queryParams *models.QueryParams) {
	var photo models.Photo
	err := c.scope.Model(&models.Photo{}).Where(&models.Photo{ID: queryParams.ID, AlbumId: queryParams.AlbumID}).Scan(&photo).Error
	if err != nil {
		RenderResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if photo.ID == 0 {
		RenderResponse(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	RenderResponse(w, http.StatusOK, photo)
}

func (c ApiController) FindAlbum(w http.ResponseWriter, queryParams *models.QueryParams) {
	var album models.Album
	err := c.scope.Model(&models.Album{}).Where(&models.Album{ID: queryParams.ID}).Scan(&album).Error
	if err != nil {
		RenderResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if album.ID == 0 {
		RenderResponse(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	RenderResponse(w, http.StatusOK, album)
}

func RenderResponse(w http.ResponseWriter, statusCode int, toWrite interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(toWrite)
}
