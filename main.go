package main

import (
	"golang-api/book"
	"golang-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
	}
	db.AutoMigrate(&book.Book{})

	// bookRepository := book.NewRepository(db)
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBook)
	router.POST("/books", bookHandler.CreateBook)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8000")
}
