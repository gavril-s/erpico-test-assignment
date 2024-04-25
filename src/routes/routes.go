package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gavril-s/erpico-test-assignemnt/service"
)

func SetupRoutes(service *service.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/gym/{id}/tt/{date}", service.GetScheduleForDate).Methods("GET")
	r.HandleFunc("/admin/gym/{id}/records/{sid}", service.GetScheduledClients).Methods("GET")

	return r
}
