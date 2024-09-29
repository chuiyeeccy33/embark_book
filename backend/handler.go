// handlers.go
package main

import (
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
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.ID = generateNextID()
	books = append(books, newBook)
	c.JSON(http.StatusOK, newBook)
}

func updateBook(c *gin.Context) {
	id, err := convertStrToInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = Book{
				ID: id, 
				Title: updatedBook.Title, 
				Author: updatedBook.Author, 
				Genre: updatedBook.Genre,
				Desc: updatedBook.Desc,
				Isbn: updatedBook.Isbn,
				Image: updatedBook.Image,
				Published: updatedBook.Published,
				Publisher: updatedBook.Publisher,
			}
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
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
