package model

import (
	"encoding/json"
	"time"
)

type myTime time.Time

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

type MusicFormRequest struct {
	ArtistName  string  `json:"artist_name" validate:"required,max=200"`
	AlbumName   string  `json:"album_name" validate:"required,max=200"`
	ImageUrl    string  `json:"image_url" validate:"max=200"`
	ReleaseDate myTime  `json:"release_date" validate:"required,max=200"`
	Price       float64 `json:"price" validate:"required,max=200"`
	SampleUrl   string  `json:"sample_url" validate:"max=200"`
}

type MusicFormResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []interface{} `json:"data,omitempty"`
}