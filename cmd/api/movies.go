package main

import (
	"greenlight.net/internal/data"
	"greenlight.net/internal/validator"
	"net/http"
)

func (app *application) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.UnprocessableEntityResponse(w, r, v.Errors)
		return
	}

	err = app.writeJson(w, http.StatusCreated, envelope{"movie": input}, nil)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}
}
