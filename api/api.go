package api

import (
	"crypto/rand"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content=type", "aplication/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "erro", err)
		SendJson(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "erro", err)
		return
	}
}

func Handler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Post("/api/shorten", HandlePost(db))
	r.Get("/{code}", HandleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"erro,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func HandlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJson(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		}

		if _, err := url.Parse(body.URL); err != nil {
			SendJson(w, Response{Error: "Invalid url passed"}, http.StatusBadGateway)
		}

		code := gencCode()
		db[code] = body.URL
		SendJson(w, Response{Data: code}, http.StatusCreated)

	}
}

func HandleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		data, ok := db[code]
		if !ok {
			http.Error(w, "Url não encontrada", http.StatusNotFound)
			http.Redirect(w, r, data, http.StatusPermanentRedirect)
		}
	}
}

const characters = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func gencCode() string {
	const n = 8

	byts := make([]byte, n)

	for i := range n {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			panic("failed to generate random number")
		}
		byts[i] = characters[nBig.Int64()]
	}

	return string(byts)
}
