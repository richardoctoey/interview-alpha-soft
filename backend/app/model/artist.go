package model

import (
	"backend/app/database"
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

func (t *Artist) ParseFromRequest(req MusicFormRequest) {
	t.ArtistName = req.ArtistName
	t.AlbumName = req.AlbumName
	t.ImageUrl = req.ImageUrl
	t.ReleaseDate = time.Time(req.ReleaseDate)
	t.Price = req.Price
	t.SampleUrl = req.SampleUrl
}

func (t *Artist) Create() error {
	return database.GetDatabase().Save(&t).Error
}

func (t *Artist) Update() error {
	return database.GetDatabase().Save(&t).Error
}