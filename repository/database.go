package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"photo-album-assignment/models"
)

const (
	databaseAddress = "143.110.190.177"
	databasePort    = "3306"
	username        = "username"
	password        = "password"
	databaseName    = "photo_album_db"
)

var (
	DBScope *gorm.DB
)

func InitDatabase() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, databaseAddress, databasePort, databaseName)
	//log.Println("Connection String ", connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln("error connecting to", err)
	}
	log.Println("connected to ", databaseName, " at ", databaseAddress, " : ", databasePort)
	DBScope = db
}

func ResetTables() {
	err := DBScope.Migrator().DropTable(&models.Album{})
	if err != nil {
		log.Println("error dropping albums ", err)
	}
	err = DBScope.Migrator().DropTable(&models.Photo{})
	if err != nil {
		log.Println("error dropping photos ", err)
	}
	err = DBScope.AutoMigrate(&models.Album{})
	if err != nil {
		log.Println("error creating table albums", err)
	}
	err = DBScope.AutoMigrate(&models.Photo{})
	if err != nil {
		log.Println("error creating table photos", err)
	}
}

func InsertData(album models.Album, photos []models.Photo) {
	scope1 := DBScope.Table("albums").Create(&album)
	scope2 := DBScope.Table("photos").Create(&photos)
	log.Println("Album: Rows Inserted = ", scope1.RowsAffected, ", Errors ( ? ) = ", scope1.Error)
	log.Println("Photos: Rows Inserted = ", scope2.RowsAffected, ", Errors ( ? ) = ", scope2.Error)
}
