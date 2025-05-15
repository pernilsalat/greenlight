package main

import (
	"fmt"
	"net/http"
)

func (app *application) LogError(req *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	body := envelope{"error": message}

	err := app.writeJson(w, status, body, nil)
	if err != nil {
		app.LogError(r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.ErrorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err map[string]string) {
	app.ErrorResponse(w, r, http.StatusUnprocessableEntity, err)
}
