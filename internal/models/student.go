package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Year      int                `json:"year" bson:"year"`
}

type StudentSearchData struct {
	TotalElements int       `json:"totalElements"`
	Data          []Student `json:"data"`
}

type StudentResponse struct {
	Status  string  `json:"status"`
	Data    Student `json:"data,omitempty"`
	Message string  `json:"message"`
}

type StudentListResponse struct {
	Status  string  `json:"status"`
	Data    Student `json:"data,omitempty"`
	Message string  `json:"message"`
}

type StudentSearchResponse struct {
	Status  string            `json:"status"`
	Data    StudentSearchData `json:"data,omitempty"`
	Message string            `json:"message"`
}
