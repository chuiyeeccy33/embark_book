// handlers.go
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func convertStrToInt(c *gin.Context, paramName string) (int, error) {
	idParam := c.Param(paramName) //get the param

	id, err := strconv.Atoi(idParam) //convert the string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return 0, err
	}
	return id, nil
}

func listBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	genre := c.PostForm("genre")
	desc := c.PostForm("desc")
	isbn := c.PostForm("isbn")
	published := c.PostForm("published")
	publisher := c.PostForm("publisher")

	// Handle the image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Save the uploaded image file
	imagePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	newBook := Book{
		ID:        len(books) + 1,
		Title:     title,
		Author:    author,
		Genre:     genre,
		Desc:      desc,
		Isbn:      isbn,
		Published: published,
		Publisher: publisher,
		Image:     imagePath, // Store the image path
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, err := convertStrToInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedBook *Book
	// Search for the existing book
	for i, book := range books {
		if book.ID == id {
			updatedBook = &books[i] // Reference the existing book for updates
			break
		}
	}

	if updatedBook == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check for optional fields and update them if present
	if title := c.PostForm("title"); title != "" {
		updatedBook.Title = title
	}
	if author := c.PostForm("author"); author != "" {
		updatedBook.Author = author
	}
	if genre := c.PostForm("genre"); genre != "" {
		updatedBook.Genre = genre
	}
	if desc := c.PostForm("desc"); desc != "" {
		updatedBook.Desc = desc
	}
	if isbn := c.PostForm("isbn"); isbn != "" {
		updatedBook.Isbn = isbn
	}
	if published := c.PostForm("published"); published != "" {
		updatedBook.Published = published
	}
	if publisher := c.PostForm("publisher"); publisher != "" {
		updatedBook.Publisher = publisher
	}

	file, err := c.FormFile("image")
	if err == nil { // File found and valid
		imagePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		updatedBook.Image = imagePath // Update the image path in the struct
	}
	c.JSON(http.StatusOK, updatedBook)
}

func deleteBook(c *gin.Context) {
	id, err := convertStrToInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func resetBooks(c *gin.Context) {
	books = nil
	FetchBooksFromAPI()
	c.JSON(http.StatusOK, books)
}
