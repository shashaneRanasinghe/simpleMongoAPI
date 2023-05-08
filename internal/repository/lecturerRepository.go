package repository

import (
	"context"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/consts"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LecturerRepository interface {
	FindAllLecturers(ctx context.Context) ([]models.Lecturer, error)
	FindLecturer(ctx context.Context, id primitive.ObjectID) (*models.Lecturer, error)
	CreateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error)
	UpdateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error)
	SearchLecturer(ctx context.Context, searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.LecturerSearchData, error)
	DeleteLecturer(ctx context.Context, id primitive.ObjectID) error
}
type lecturerRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewLecturerRepo(client *mongo.Client, collection *mongo.Collection) *lecturerRepo {
	return &lecturerRepo{
		client:     client,
		collection: collection,
	}
}

func (l lecturerRepo) FindAllLecturers(ctx context.Context) ([]models.Lecturer, error) {
	var results []models.Lecturer

	cursor, err := l.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Error(consts.DBCursorCloseError)
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &results); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info(consts.LogLecturers, results)
	return results, nil
}

func (l lecturerRepo) FindLecturer(ctx context.Context, id primitive.ObjectID) (*models.Lecturer, error) {
	var result models.Lecturer

	err := l.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(consts.LogLecturers, result)
	return &result, nil
}

func (l lecturerRepo) CreateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error) {

	lecturer.ID = primitive.NewObjectID()

	result, err := l.collection.InsertOne(ctx, lecturer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(consts.LogLecturers, result.InsertedID)
	lecturer.ID = result.InsertedID.(primitive.ObjectID)
	return lecturer, nil
}

func (l lecturerRepo) UpdateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error) {

	_, err := l.collection.ReplaceOne(ctx, bson.M{"_id": lecturer.ID}, lecturer)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(consts.LogLecturers, lecturer)
	return lecturer, nil
}

func (l lecturerRepo) SearchLecturer(ctx context.Context, searchString string, pagination models.Pagination, sortBy models.SortBy) (*models.LecturerSearchData, error) {
	sortDirection := consts.SortASC
	if sortBy.Direction == consts.DESC {
		sortDirection = consts.SortDESC
	}
	var results []models.Lecturer

	sortOptions := options.Find().SetSort(bson.M{sortBy.Column: sortDirection})
	paginationOptions := options.Find().SetSkip(int64((pagination.Page - 1) * pagination.PageSize)).
		SetLimit(int64(pagination.PageSize))

	cursor, err := l.collection.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{
					"firstname": bson.M{
						"$regex":   searchString,
						"$options": "i",
					},
				},
				bson.M{
					"lastname": bson.M{
						"$regex":   searchString,
						"$options": "i",
					},
				},
			},
		},
		sortOptions, paginationOptions)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Error(consts.DBCursorCloseError)
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &results); err != nil {
		log.Error(err)
		return nil, err
	}

	searchResp := models.LecturerSearchData{
		TotalElements: 0,
		Data:          results,
	}
	log.Info(consts.LogLecturers, results)
	return &searchResp, nil
}

func (l lecturerRepo) DeleteLecturer(ctx context.Context, id primitive.ObjectID) error {

	result, err := l.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Record Deleted Successfully", result.DeletedCount)
	return nil
}
