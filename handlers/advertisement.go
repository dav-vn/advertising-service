package handlers

import (
	"encoding/json"
	"github.com/dav-vn/advertising-service/models"
	"net/http"
	"strconv"
)

type AdvertisementHandler struct {
	Storage *models.AdvertisementStorage
}

func (h *AdvertisementHandler) GetAdvertisements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	advertisement, err := h.Storage.GetAdvertisementByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(advertisement)
}

func (h *AdvertisementHandler) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ad models.Advertisement
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newAd, err := h.Storage.CreateAdvertisement(ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newAd)
}
