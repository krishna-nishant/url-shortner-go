package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"urlshortener/api/handlers"
	"urlshortener/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a Gin router
	router := gin.Default()

	// Initialize our store with persistence
	store := db.NewMemoryStore(true)

	// Create URL handler
	urlHandler := &handlers.URLHandler{Store: store}

	// Serve static files
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Routes
	router.GET("/", urlHandler.HomePage)
	router.POST("/shorten", urlHandler.CreateShortURL)
	router.GET("/:shortURL", urlHandler.RedirectToOriginal)
	router.GET("/api/urls", urlHandler.GetAllURLs)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Println("Starting server on port", port)
	router.Run(":" + port)
}
