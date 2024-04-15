package routes

import (
	"net/http"
	"strconv"

	"github.com/JoaquinBadillo/FlightsAPI/db/models"
	provider "github.com/JoaquinBadillo/FlightsAPI/db/provider"
	"github.com/JoaquinBadillo/FlightsAPI/lib"
)

type SeatsResponse struct {
	Seats []*models.Seat `json:"seats"`
}

func GetSeats(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid flight id", http.StatusBadRequest)
		return
	}

	seats, err := provider.Mgr.GetAvailableSeats(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &SeatsResponse{
		Seats: seats,
	}

	lib.WriteResponse(response, w, http.StatusOK)
}
