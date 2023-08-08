package controllers

import (
	"encoding/json"
	"net/http"
	"urlgo/services"
)

type URLController struct {
	URLService services.URLService
}

func NewUrlController(urlService services.URLService) *URLController {
	return &URLController{URLService: urlService}
}

func (c *URLController) CreateUrl(w http.ResponseWriter, r *http.Request) {
	original := r.FormValue("original")
	if original == "" {
		http.Error(w, "original URL is required", http.StatusBadRequest)

		return
	}

	url, err := c.URLService.CreateUrl(original)
	if err != nil {
		http.Error(w, "Failed to create shortUrl", http.StatusInternalServerError)

		return
	}

	response, _ := json.Marshal(url)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *URLController) GetUrlByShortCode(w http.ResponseWriter, r *http.Request) {
	shortCode := r.FormValue("shortCode")
	if shortCode == "" {
		http.Error(w, "shortCode is required to get URL", http.StatusBadRequest)

		return
	}

	url, err := c.URLService.GetUrlByShortCode(shortCode)
	if err != nil {
		http.Error(w, "ShortCode not exists in DB", http.StatusBadRequest)

		return
	}

	response, _ := json.Marshal(url)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
