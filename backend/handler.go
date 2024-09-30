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
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, err := convertStrToInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			if title, ok := updatedFields["title"]; ok {
				books[i].Title = title.(string)
			}
			if author, ok := updatedFields["author"]; ok {
				books[i].Author = author.(string)
			}
			if genre, ok := updatedFields["genre"]; ok {
				books[i].Genre = genre.(string)
			}
			if desc, ok := updatedFields["desc"]; ok {
				books[i].Desc = desc.(string)
			}
			if isbn, ok := updatedFields["isbn"]; ok {
				books[i].Isbn = isbn.(string)
			}
			if image, ok := updatedFields["image"]; ok {
				books[i].Image = image.(string)
			}
			if published, ok := updatedFields["published"]; ok {
				books[i].Published = published.(string)
			}
			if publisher, ok := updatedFields["publisher"]; ok {
				books[i].Publisher = publisher.(string)
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
