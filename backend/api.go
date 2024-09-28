package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Book struct {
	ID        int    `json: id`
	Title     string `json: title`
	Author    string `json: author`
	Genre     string `json: genre`
	Desc      string `json: genre`
	Isbn      string `json: isbn`
	Image     string `json: image`
	Published string `json: published`
	Publisher string `json: publisher`
}

type Response struct {
	Data []Book `json: data`
}

var ori []Book
var current []Book

func Fetch() ([]Book, error) {
	response, err := http.Get("https://fakerapi.it/api/v1/books")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch books")
	}
	var bookReponse Response
	if err := json.NewDecoder((response.Body).Decode(&bookResponse); err != nil {
		return nil, err
	})

	ori = bookReponse.Data
	current = ori
	return current, nil
}

func Create(newBook Book) {
	current = append(current, newBook)
}

func Update(updateBook Book) error {
	for i, book := range current {
		if book.ID == updateBook.ID {
			current[i] = updateBook
			return nil
		}
	}
	return errors.New("not found")
}

func Delete(id int) error {
	for i, book := range current {
		if book.ID == id {
			current = append(current[:i], current[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func Reset() {
	current = ori
}
