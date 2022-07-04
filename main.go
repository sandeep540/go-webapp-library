package main

import (
	"library/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	book := new(controller.BookController)

	router := gin.Default()
	router.GET("/books", book.GetBooks)
	router.POST("/book", book.PostBook)

	router.Run("localhost:8080")

}
