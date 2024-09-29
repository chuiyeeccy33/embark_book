//main.go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
	// initialize the book data
	FetchBooksFromAPI()

    r := gin.Default()

    // API routes
    r.GET("/books", listBooks)
    r.POST("/books", createBook)
    r.PUT("/books/:id", updateBook)
    r.DELETE("/books/:id", deleteBook)
    r.POST("/books/reset", resetBooks)

    // Serve static files (frontend)
    // r.Static("/static", "./static")

    // Run the server
    r.Run(":8080") // Listen on port 8080
}
