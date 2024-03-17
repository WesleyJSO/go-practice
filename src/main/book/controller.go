package book

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := AddCase(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	jBook, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jBook)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var requestBook Book
	if err := json.NewDecoder(r.Body).Decode(&requestBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateBook, err := FindCase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jBook, err := json.Marshal(requestBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateBook.Pages = requestBook.Pages
	updateBook.Title = requestBook.Title
	//TODO update it
	w.WriteHeader(http.StatusCreated)
	w.Write(jBook)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	deletedBook, err := DeleteCase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jBook, err := json.Marshal(deletedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jBook)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := FindCase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jBook, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jBook)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	books := FindAllCase()
	jBooks, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jBooks)
}
