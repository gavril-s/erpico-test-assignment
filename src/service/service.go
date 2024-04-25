package service

import (
    "time"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gavril-s/erpico-test-assignemnt/repository"
)

type Service struct {
	Repo repository.Repository
}

func (s *Service) GetScheduleForDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gymID := params["id"]
	dateStr := params["date"]

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	schedule, err := s.Repo.GetScheduleForDate(gymID, date)
	if err != nil {
		http.Error(w, "Failed to fetch schedule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}

func (s *Service) GetScheduledClients(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gymID := params["id"]
	scheduleID := params["sid"]

	clients, err := s.Repo.GetScheduledUsers(gymID, scheduleID)
	if err != nil {
		http.Error(w, "Failed to fetch scheduled clients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
