package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/cassianobraz/SearchForMovieInformation/omdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handlerSearchMovie(apiKey))

	return r
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer mashal de json", "error", err)
		sendJSON(w, Response{Error: "something went wrong "}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error ao enviar a resposta", "error", err)
		return
	}
}

func handlerSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)
		if err != nil {
			sendJSON(w, Response{Error: "something wrong with omdb"}, http.StatusBadGateway)
			return
		}

		sendJSON(w, Response{Data: res}, http.StatusOK)
	}
}
