package models

import (
	"github.com/Shivam5567/golang/Projects/pkg/config"
	"gorm.io/gorm"
)

var db*gorm.DB
type Book struct{
	gorm.Model
	Name string `gorm"":json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`	
}

func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book)CreateBook()*Book{
	db.Create(&b)
	return b
}
func GetAllBooks()[]Book{
	var Books[]Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64)(*Book,*gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

// In models/book.go

func DeleteBook(ID int64) *gorm.DB {
	dbResult := db.Unscoped().Where("ID=?", ID).Delete(&Book{})
	return dbResult
}