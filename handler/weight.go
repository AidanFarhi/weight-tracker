package handler

import (
	"encoding/json"
	"net/http"
	"weight-tracker/middleware"
	"weight-tracker/service"
)

type WeightHandler struct {
	ws *service.WeightService
}

func NewWeightHandler(ws *service.WeightService) *WeightHandler {
	return &WeightHandler{ws}
}

func (wh *WeightHandler) GetDailyWeights(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	weightEntries, err := wh.ws.GetDailyWeightEntriesForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(weightEntries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
