package controller

import (
	"net/http"
	"weight-tracker/middleware"
	"weight-tracker/model"
	"weight-tracker/service"
)

type HomeController struct {
	ws *service.WeightService
}

func NewHomeController(ws *service.WeightService) *HomeController {
	return &HomeController{ws}
}

func (hc *HomeController) GetHome(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	latestWeightEntry, err := hc.ws.GetLatestDailyWeightEntryForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	latestTargetWeight, err := hc.ws.GetLatestTargetWeightForUser(userId)
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
