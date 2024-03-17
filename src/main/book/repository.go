package book

import (
	"errors"
	"fmt"
)

var books = make(map[int]*Book)
var keys = make([]int, 0, 10)

func AddBook(book *Book) (int, error) {
	var id int
	if len(keys) == 0 {
		keys = append(keys, 1)
		id = 1
	} else {
		id = keys[len(keys)-1] + 1
	}
	books[id] = book
	return id, nil
}

func DeleteBook(id int) {
	delete(books, id)
}

func FindBook(id int) (*Book, error) {
	book := books[id]
	if book == nil {
		msg := fmt.Sprintf("book not found for id: %d", id)
		return nil, errors.New(msg)
	}
	return book, nil
}

func FindBooks() map[int]*Book {
	return books
}
