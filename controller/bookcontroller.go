package controller

import (
	"context"
	"fmt"
	"library/config"
	"library/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type BookController struct{}

var booksCollection *mongo.Collection = config.GetCollection(config.DB, "books")

func (b BookController) GetBooks(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var books []model.Book
	defer cancel()

	results, err := booksCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var book model.Book
		if err = results.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		books = append(books, book)
	}

	c.JSON(http.StatusOK,
		UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": books}},
	)
}

func (b BookController) PostBook(ctx *gin.Context) {

	fmt.Println("PostBook")
	var newBook model.Book

	// Call BindJSON to bind the received JSON to
	if err := ctx.BindJSON(&newBook); err != nil {
		fmt.Println("Error:", err)
		return
	}

	book := model.Book{
		Id:    primitive.NewObjectID(),
		Title: newBook.Title,
		Year:  newBook.Year,
	}

	result, err := booksCollection.InsertOne(ctx, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusCreated, UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

}
