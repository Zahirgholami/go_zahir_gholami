package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zahirgholami/go_zahir_gholami/internal/models"
	"github.com/Zahirgholami/go_zahir_gholami/internal/services"
)

func GetUnits(unitService services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		units, err := unitService.GetUnits()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonUnits, err := json.Marshal(units)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUnits)
	}
}

func CreateUnit(unitService services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unit models.Unit
		err := json.NewDecoder(r.Body).Decode(&unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdUnit, err := unitService.CreateUnit(unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonUnit, err := json.Marshal(createdUnit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonUnit)
	}
}
