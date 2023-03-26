package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hammad-umar/go-movies-api/pkg/models"
	"github.com/hammad-umar/go-movies-api/pkg/utils"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := models.GetAllMovies()
	res, _ := json.Marshal(movies)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing bookId: ", err)
	}

	movie, _ := models.GetMovieDetails(id)
	res, _ := json.Marshal(movie)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	createMovie := &models.Movie{}
	utils.ParseBody(r, createMovie)
	movie := createMovie.CreateMovie()

	res, _ := json.Marshal(movie)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing bookId: ", err)
	}

	m := models.DeleteMovie(id)
	res, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing bookId: ", err)
	}

	updateMovie := &models.Movie{}
	utils.ParseBody(r, updateMovie)
	m, db := models.GetMovieDetails(id)

	if updateMovie.Title != "" {
		m.Title = updateMovie.Title
	}

	if updateMovie.Description != "" {
		m.Description = updateMovie.Description
	}

	if updateMovie.Rating != 0 {
		m.Rating = updateMovie.Rating
	}

	if updateMovie.Year != 0 {
		m.Year = updateMovie.Year
	}

	db.Save(&m)
	res, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
