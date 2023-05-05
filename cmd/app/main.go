package main

import (
	"context"
	"fmt"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/config"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/repository"
	student2 "github.com/shashaneRanasinghe/simpleMongoAPI/internal/usecases/student"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/database"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	config.LoadConfigs()
	client := database.IntiDB()
	collection := client.Database("SchoolDB").Collection("students")

	studentRepo := repository.NewStudentRepo(client, collection)
	student := student2.NewStudent(studentRepo)
	ctx := context.Background()
	id, _ := primitive.ObjectIDFromHex("6454d313546f9599d7a469dd")

	stupd := models.Student{
		ID:        id,
		FirstName: "Max",
		LastName:  "Verstappen",
		Year:      2,
	}
	st := models.Student{
		ID:        primitive.NewObjectID(),
		FirstName: "Max",
		LastName:  "Verstappen",
		Year:      2,
	}
	pagination := models.Pagination{
		Page:     0,
		PageSize: 10,
	}
	sort := models.SortBy{
		Column:    "firstname",
		Direction: "ASC",
	}

	allst, err := student.GetAllStudents(ctx)
	fmt.Printf("get all  %v  %v ", allst, err)
	onest, err := student.GetStudent(ctx, "6454d313546f9599d7a469dd")
	fmt.Printf("get 1  %v  %v ", onest, err)
	crst, err := student.CreateStudent(ctx, &st)
	fmt.Printf("create  %v  %v ", crst, err)
	upst, err := student.UpdateStudent(ctx, &stupd)
	fmt.Printf("Update  %v  %v ", upst, err)
	searchst, err := student.SearchStudent(ctx, "x", pagination, sort)
	fmt.Printf("Search %v  %v   ", searchst, err)
	err = student.DeleteStudent(ctx, "6454d313546f9599d7a469dd")
	fmt.Printf("Delete %s ", err)
	database.DisconnectDB(client)

	log.Info("Server Stopped")
}
