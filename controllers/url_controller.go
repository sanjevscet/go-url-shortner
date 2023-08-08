package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"urlgo/services"
)

type URLController struct {
	URLService services.URLService
}

func NewUrlController(urlService services.URLService) *URLController {
	return &URLController{URLService: urlService}
}

type CreateRequest struct {
	Original string `json:"original"`
}

type GetRequest struct {
	ShortCode string `json:"shortCode"`
}

func (c *URLController) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var createRequest CreateRequest

	// read the data from request body
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	log.Println(err)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)

		return
	}

	if createRequest.Original == "" {
		http.Error(w, "original URL is required", http.StatusBadRequest)

		return
	}

	url, err := c.URLService.CreateUrl(createRequest.Original)
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
	// get all query parameters
	queryParams := r.URL.Query()

	//retrieves values for code key
	code := queryParams.Get("code")

	if len(code) == 0 {
		http.Error(w, "shortCode is required to get URL", http.StatusBadRequest)

		return
	}

	url, err := c.URLService.GetUrlByShortCode(code)
	if err != nil {
		http.Error(w, "ShortCode not exists in DB", http.StatusBadRequest)

		return
	}

	response, _ := json.Marshal(url)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
