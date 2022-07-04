package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Title string             `json:"title"`
	Year  string             `json:"year"`
}
