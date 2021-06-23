package models

type QueryParams struct {
	Type    string `schema:"type"`
	ID      int    `schema:"id"`
	AlbumID int    `schema:"album"`
}
