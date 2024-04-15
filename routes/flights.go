/* Flights

Contains the API routes for the flights resource.
- GetFlights[Paginated]: Get all the available flights with optional filters.
- GetFlight: Get a specific flight by id.

Joaquin Badillo
2024-04-14
*/

package routes

import (
	"net/http"
	"strconv"

	"github.com/JoaquinBadillo/FlightsAPI/db/models"
	provider "github.com/JoaquinBadillo/FlightsAPI/db/provider"
	"github.com/JoaquinBadillo/FlightsAPI/lib"
)

type FlightsResponse struct {
	Flights []*models.Flight `json:"flights"`
}

func GetFlights(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		http.Error(w, "Invalid page size", http.StatusBadRequest)
		return
	}

	var flights []*models.Flight
	city := r.URL.Query().Get("city")
	country := r.URL.Query().Get("country")

	if city == "" || country == "" {
		flights, err = provider.Mgr.GetAvailableFlights(limit, offset-1)
	} else {
		flights, err = provider.Mgr.GetAvailableFlightsByLocation(city, country, limit, offset-1)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &FlightsResponse{
		Flights: flights,
	}

	lib.WriteResponse(response, w, http.StatusOK)
}

func GetFlight(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid flight id", http.StatusBadRequest)
		return
	}

	flight, err := provider.Mgr.GetFlight(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lib.WriteResponse(flight, w, http.StatusOK)
}
