package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// dummy data from faker api
type Book struct {
	ID        int    `json: "id"`
	Title     string `json: "title"`
	Author    string `json: "author"`
	Genre     string `json: "genre"`
	Desc      string `json: "description"`
	Isbn      string `json: "isbn"`
	Image     string `json: "image"`
	Published string `json: "published"`
	Publisher string `json: "publisher"`
}

var books []Book

func FetchBooksFromAPI() {
	// Fetch data from faker API
	response, err := http.Get("https://fakerapi.it/api/v1/books")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal directly into the internal Book structure
	var result struct {
		Data []Book `json:"data"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	books = result.Data // add api result into books array

	fmt.Println("Fetched and populated books:", books)
}

func generateNextID() int {
	if len(books) == 0 {
		// Start with ID "1" if there are no books
		return 1
	}

	// Find the last book's ID and convert it to an integer
	lastBook := books[len(books)-1]
	return lastBook.ID + 1
}
