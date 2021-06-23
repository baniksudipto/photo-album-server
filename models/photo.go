package models


type Photo struct {
	AlbumId      int    `json:"albumId" gorm:"index:photos_idx,priority:1"`
	ID           int    `json:"id" gorm:"index:photos_idx,priority:2"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}