// main.go
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the book data
	FetchBooksFromAPI()

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},                 // Update this with your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, // Include PATCH here
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // Allow credentials
	}))

	// API routes
	r.GET("/api/books", listBooks)
	r.POST("/api/books", createBook)
	r.PATCH("/api/books/:id", updateBook)
	r.DELETE("/api/books/:id", deleteBook)
	r.POST("/api/books/reset", resetBooks)

	// Serve static files (frontend)
	r.Static("/static", "./frontend/dist")

	// Run the server
	r.Run(":8000") // Listen on port 8000
}
