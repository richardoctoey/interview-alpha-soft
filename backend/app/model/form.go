package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type myTime time.Time

func (mt *myTime) String() string {
	x := time.Time(*mt)
	return x.Format("2006-01-02")
}

func (t *myTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = myTime(v)
	case nil:
		*t = myTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

func (t myTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format("2006-01-02")), nil
}

func (t *myTime) UnmarshalText(value string) error {
	dd, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	*t = myTime(dd)
	return nil
}

func (myTime) GormDataType() string {
	return "DATE"
}

func (t myTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02"))
}

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
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data,omitempty"`
}
