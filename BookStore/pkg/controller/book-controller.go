package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shivam5567/golang/Projects/pkg/models"
	"github.com/Shivam5567/golang/Projects/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter,r*http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("err while parsing")
	}
	bookDetails,_:=models.GetBookById(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r * http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_:=json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r* http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book :=models.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Parse the incoming JSON into an update struct
    var updateBook = &models.Book{}
    utils.ParseBody(r, updateBook)

    // Extract bookId from URL
    vars := mux.Vars(r)
    bookId := vars["bookId"]

    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("Error while parsing book ID:", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Fetch existing book
    bookDetails, db := models.GetBookById(ID)
    if bookDetails == nil {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"error": "Book not found"}`))
        return
    }

    // Update fields only if provided
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }

    // Save updated book
    db.Save(&bookDetails)

    // Respond with JSON
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
