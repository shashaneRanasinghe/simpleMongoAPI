package resolvers

import (
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/repository"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/usecases/student"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	studentUsecase student.StudentUsecase
}

func NewResolver(client *mongo.Client) *Resolver {
	studentDBCollection := client.Database("SchoolDB").Collection("students")
	studentRepo := repository.NewStudentRepo(client, studentDBCollection)
	studentUsecase := student.NewStudent(studentRepo)
	return &Resolver{
		studentUsecase: studentUsecase,
	}
}
