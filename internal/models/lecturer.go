package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lecturer struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Year      int                `json:"year" bson:"year"`
}

type LecturerSearchData struct {
	TotalElements int        `json:"totalElements"`
	Data          []Lecturer `json:"data"`
}

type LecturerResponse struct {
	Status  string   `json:"status"`
	Data    Lecturer `json:"data,omitempty"`
	Message string   `json:"message"`
}

type LecturerListResponse struct {
	Status  string   `json:"status"`
	Data    Lecturer `json:"data,omitempty"`
	Message string   `json:"message"`
}

type LecturerSearchResponse struct {
	Status  string             `json:"status"`
	Data    LecturerSearchData `json:"data,omitempty"`
	Message string             `json:"message"`
}
