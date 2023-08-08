package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"urlgo/models"
)

type URLService interface {
	CreateUrl(original string) (*models.URL, error)
	GetUrlByShortCode(shortCode string) (*models.URL, error)
}

type urlService struct {
	db *sql.DB
}

func NewUrlService(db *sql.DB) URLService {

	return &urlService{db: db}
}

func (s *urlService) CreateUrl(original string) (*models.URL, error) {
	shortCode := generateShortCode(6)
	query := "INSERT into urls (original, shortCode) VALUES (?, ?)"
	result, err := s.db.Exec(query, original, shortCode)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to stored url in DB")
	}

	lastInsertedId, _ := result.LastInsertId()

	return &models.URL{
		ID:        uint(lastInsertedId),
		Original:  original,
		ShortCode: shortCode,
	}, nil
}

func (s *urlService) GetUrlByShortCode(shortCode string) (*models.URL, error) {
	var url models.URL
	query := "SELECT id, original from urls where shortCode = ?"
	err := s.db.QueryRow(query, shortCode).Scan(&url.ID, &url.Original)
	if err != nil {
		log.Println(err)

		if err == sql.ErrNoRows {
			return nil, errors.New("short url not found")
		}

		return nil, errors.New("failed to fetch url from DB")
	}

	url.ShortCode = shortCode

	return &url, nil
}

func generateShortCode(length int) string {
	chartSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charSetBytes := []byte(chartSet)

	// create a buffer to hold random bytes
	randomBytes := make([]byte, length)

	// read random bytes from crypto/rand
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	//convert random bytes to indexes in the character set
	for i := 0; i < length; i++ {
		randomBytes[i] = charSetBytes[randomBytes[i]%byte(len(charSetBytes))]
	}

	// encode the bytes to string by using base64 encoding
	shortCode := base64.RawURLEncoding.EncodeToString(randomBytes)

	return shortCode[:length]
}
