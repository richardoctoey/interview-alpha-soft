package model

import (
	"backend/app/database"
	"html"
	"time"
)

type Artist struct {
	ArtistId int `gorm:"primary_key;column:artist_id"`
	ArtistName string
	AlbumName string
	ImageUrl string
	ReleaseDate time.Time
	Price float64
	SampleUrl string
}

func (t *Artist) TableName() string {
	return "artist"
}

func (t *Artist) ParseFromRequest(req MusicFormRequest) {
	t.ArtistName = html.EscapeString(req.ArtistName)
	t.AlbumName = html.EscapeString(req.AlbumName)
	t.ImageUrl = html.EscapeString(req.ImageUrl)
	t.ReleaseDate = time.Time(req.ReleaseDate)
	t.Price = req.Price
	t.SampleUrl = html.EscapeString(req.SampleUrl)
}

func (t *Artist) Create() error {
	return database.GetDatabase().Create(&t).Error
}

func (t *Artist) Update() error {
	return database.GetDatabase().Save(&t).Error
}

func (t *Artist) FindArtistById(id string) error {
	db := database.GetDatabase()
	err := db.Where("artist_id = ?", id).First(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func ArtistDeleteById(id string) error {
	db := database.GetDatabase()
	obj := Artist{}
	err := obj.FindArtistById(id)
	if err != nil {
		return err
	}
	return db.Delete(&obj).Error
}