package models

import (
	"github.com/hammad-umar/go-movies-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	ID 			uint		`gorm:"primaryKey" json:"id"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description"`
	Rating 		float64		`json:"rating"`
	Year 		int64		`json:"year"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func GetAllMovies() []Movie {
	var movies []Movie 
	db.Find(&movies)
	return movies 
}

func GetMovieDetails(id int64) (*Movie, *gorm.DB) {
	var movie Movie 
	db := db.Where("id=?", id).Find(&movie)
	return &movie, db 
}

func DeleteMovie(id int64) *Movie {
	var movie Movie 
	db.Where("id=?", id).Delete(&movie)
	return &movie 
}

func (m *Movie) CreateMovie() *Movie {
	db.Create(&m)
	return m 
}
