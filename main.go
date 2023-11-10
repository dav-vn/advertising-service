package main

import (
	"log"
	"net/http"

	"github.com/dav-vn/advertising-service/handlers"
	"github.com/dav-vn/advertising-service/models"
)

func main() {
	// Инициализация хранилища
	storage := &models.AdvertisementStorage{
		Advertisements: make([]models.Advertisement, 0),
	}

	// Инициализация обработчика
	handler := &handlers.AdvertisementHandler{
		Storage: storage,
	}

	// Роутинг API
	http.HandleFunc("/advertisements", handler.GetAdvertisements)
	http.HandleFunc("/advertisement", handler.GetAdvertisement)
	http.HandleFunc("/advertisement/create", handler.CreateAdvertisement)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
