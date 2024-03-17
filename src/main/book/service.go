package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"library/src/main/producer"
)

func FindCase(id int) (*Book, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	return FindBook(id)
}

func FindAllCase() map[int]*Book {
	return FindBooks()
}

func AddCase(book *Book) error {
	if book == nil {
		return errors.New("book pointer doesn't reference a valid struct")
	}
	if err := validate(book); err != nil {
		return err
	}
	if id, err := AddBook(book); err != nil {
		return err
	} else {
		book.Id = id
	}
	jBook, err := json.Marshal(book)
	if err != nil {
		return err
	}
	producer.SendMsg("library-books", jBook)
	return nil
}

func DeleteCase(id int) (*Book, error) {
	book, err := FindBook(id)
	if err != nil {
		return nil, err
	}
	DeleteBook(book.Id)
	return book, nil
}

func validate(book *Book) error {
	if book.Pages < 0 {
		return fmt.Errorf("invalid page number informed %d, pages should be greater than 0", book.Pages)
	}
	if len(book.Title) == 0 {
		return fmt.Errorf("invalid title informed %q, title len must be greater than 0", book.Title)
	}
	return nil
}
