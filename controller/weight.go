package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (wh *WeightController) PostDailyWeightEntry(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	weightStr := r.FormValue("weight")
	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// make a call to the service to put a weight entry value
	fmt.Println(userId, weight)
	http.Redirect(w, r, "/", http.StatusFound)
}
