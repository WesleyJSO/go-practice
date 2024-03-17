package test

import (
	"library/src/main/book"
	"testing"
)

const wrongErrMsg string = "unexpected error: %q, error message should be: %q"

func TestDeleteBookInvalidId(t *testing.T) {
	var errMsg string = "book not found for id: 1"

	_, err := book.DeleteCase(1)

	if err.Error() != errMsg {
		t.Errorf(wrongErrMsg, err, errMsg)
	}
}

func TestDeleteValidBook(t *testing.T) {
	var b *book.Book = &book.Book{Title: "test book", Pages: 123}

	book.AddCase(b)

	_, err := book.DeleteCase(b.Id)

	if err != nil {
		t.Errorf("unexpected error: %q, book with id: %d should have been successfully deleted", err, b.Id)
	}
}

func TestFindInvalidBookId(t *testing.T) {
	var id int = 0
	var errMsg string = "invalid id"

	_, err := book.FindCase(id)

	if err.Error() != errMsg {
		t.Errorf(wrongErrMsg, err, errMsg)
	}
}

func TestFindAllBooks(t *testing.T) {
	var books map[int]*book.Book = book.FindAllCase()

	if len(books) > 0 {
		t.Errorf("unexpected books map, it should be empty, found %v", books)
	}
}

func TestAddValidBook(t *testing.T) {
	var b *book.Book = &book.Book{Title: "test book", Pages: 123}

	err := book.AddCase(b)

	if err != nil {
		t.Errorf("unexpected error %q, err should be nil", err)
	}
}

func TestAddNilBook(t *testing.T) {
	var b *book.Book = nil
	var errMsg string = "book pointer doesn't reference a valid struct"

	err := book.AddCase(b)

	if err.Error() != errMsg {
		t.Errorf(wrongErrMsg, err, errMsg)
	}
}

func TestValidBookAmountOfPages(t *testing.T) {
	var b *book.Book = &book.Book{Title: "test title", Pages: -1}
	var errMsg = "invalid page number informed -1, pages should be greater than 0"

	err := book.AddCase(b)

	if err.Error() != errMsg {
		t.Errorf(wrongErrMsg, err, errMsg)
	}
}

func TestValidBookTitle(t *testing.T) {
	var b *book.Book = &book.Book{Title: "", Pages: 0}
	var errMsg = "invalid title informed \"\", title len must be greater than 0"

	err := book.AddCase(b)

	if err.Error() != errMsg {
		t.Errorf(wrongErrMsg, err, errMsg)
	}
}
