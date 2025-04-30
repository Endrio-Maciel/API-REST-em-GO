package api

import (
	service "api_rest/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *Application) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := app.Service.FindAll()
	SendJson(w, Response{Data: users}, http.StatusOK)
}

func (app *Application) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	uid, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "invalid UUID"}, http.StatusBadRequest)
		return
	}

	user, ok := app.Service.FindById(service.Id(uid))
	if !ok {
		SendJson(w, Response{Error: "user not found"}, http.StatusNotFound)
	}

	SendJson(w, Response{Data: user}, http.StatusOK)

}
func (app *Application) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var body service.PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		SendJson(w, Response{Error: "invalid body"}, http.StatusBadRequest)
		return
	}
	user := app.Service.Insert(body)
	SendJson(w, Response{Data: user}, http.StatusCreated)
}

func (app *Application) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	uid, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "invalid UUID"}, http.StatusBadRequest)
		return
	}

	var body service.PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		SendJson(w, Response{Error: "invalid body"}, http.StatusBadRequest)
		return
	}

	updated, ok := app.Service.Update(service.Id(uid), body)
	if !ok {
		SendJson(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: updated}, http.StatusOK)
}

func (app *Application) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	uid, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "invalid UUID"}, http.StatusBadRequest)
		return
	}

	ok := app.Service.Delete(service.Id(uid))
	if !ok {
		SendJson(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: "deleted"}, http.StatusOK)
}
