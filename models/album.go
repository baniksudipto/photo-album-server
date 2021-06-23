package models

type Album struct {
	ID     int  `json:"id" gorm:"index:album_pkey,unique"`
	UserId int  `json:"userId"`
	Title  string `json:"title"`
}
