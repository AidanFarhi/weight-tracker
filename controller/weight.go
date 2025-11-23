package controller

import (
	"encoding/json"
	"net/http"
	"weight-tracker/middleware"
	"weight-tracker/service"
)

type WeightController struct {
	ws *service.WeightService
}

func NewWeightController(ws *service.WeightService) *WeightController {
	return &WeightController{ws}
}

func (wh *WeightController) GetDailyWeights(w http.ResponseWriter, r *http.Request) {
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

func (wh *WeightController) GetDailyWeightEntry(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "daily-weight-entry", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
