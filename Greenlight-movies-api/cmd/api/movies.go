package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) createNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]any{
		"message": "Created successfully",
		"movie":   "...",
	})
	if err != nil {
		app.logger.Error("Error encoding post movies response", err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	movieId, err := strconv.Atoi(params.ByName("id"))
	if err != nil || movieId < 1 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"movies": fmt.Sprintf("Movie detail of %d", movieId),
	})
	if err != nil {
		app.logger.Error("Error encoding get movie response", err)
		http.Error(w, "Error getting movie ", http.StatusInternalServerError)
	}
}
