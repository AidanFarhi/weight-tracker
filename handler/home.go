package handler

import (
	"net/http"
	"weight-tracker/middleware"
	"weight-tracker/model"
	"weight-tracker/service"
)

type HomeHandler struct {
	ws *service.WeightService
}

func NewHomeHandler(ws *service.WeightService) *HomeHandler {
	return &HomeHandler{ws}
}

func (hh *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	latestWeightEntry, err := hh.ws.GetLatestDailyWeightEntryForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	latestTargetWeight, err := hh.ws.GetLatestTargetWeightForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = RenderPage(w, "home", model.HomePageData{
		CurrentWeight: latestWeightEntry.Weight,
		TargetWeight:  latestTargetWeight.Weight,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
