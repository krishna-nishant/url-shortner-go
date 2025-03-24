package handlers

import (
	"math/rand"
	"net/http"

	"urlshortener/db"

	"github.com/gin-gonic/gin"
)

// URLHandler handles URL-related HTTP requests
type URLHandler struct {
	Store *db.MemoryStore
}

// HomePage renders the home page
func (h *URLHandler) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "URL Shortener",
	})
}

// CreateShortURL creates a shortened URL
func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var urlRequest struct {
		Original string `json:"original" binding:"required"`
	}

	if err := c.ShouldBindJSON(&urlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a short URL
	shortCode := generateShortURL()

	// Create URL object
	newURL := db.URL{
		Original: urlRequest.Original,
		Short:    shortCode,
	}

	// Save to store
	savedURL := h.Store.SaveURL(newURL)

	c.JSON(http.StatusOK, savedURL)
}

// RedirectToOriginal redirects from a short URL to the original URL
func (h *URLHandler) RedirectToOriginal(c *gin.Context) {
	shortURL := c.Param("shortURL")

	url, exists := h.Store.GetByShortURL(shortURL)
	if !exists {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	// Update click count
	h.Store.UpdateClickCount(shortURL)

	// Redirect to original URL
	c.Redirect(http.StatusMovedPermanently, url.Original)
}

// GetAllURLs returns all shortened URLs
func (h *URLHandler) GetAllURLs(c *gin.Context) {
	urlList := h.Store.GetAllURLs()
	c.JSON(http.StatusOK, urlList)
}

// generateShortURL generates a random short URL
func generateShortURL() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""
	for i := 0; i < 6; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}
	return result
}
