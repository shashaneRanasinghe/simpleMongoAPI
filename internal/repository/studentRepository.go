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

type StudentRepository interface {
	FindAllStudents(ctx context.Context) ([]models.Student, error)
	FindStudent(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	SearchStudent(ctx context.Context, searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.StudentSearchData, error)
	DeleteStudent(ctx context.Context, id string) error
}

type studentRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStudentRepo(client *mongo.Client, collection *mongo.Collection) *studentRepo {
	return &studentRepo{
		client:     client,
		collection: collection,
	}
}

func (s *studentRepo) FindAllStudents(ctx context.Context) ([]models.Student, error) {
	var results []models.Student

	cursor, err := s.collection.Find(ctx, bson.M{})
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
	log.Info("Students : ", results)
	return results, nil
}

func (s *studentRepo) FindStudent(ctx context.Context, id string) (*models.Student, error) {
	var result models.Student

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(consts.ObjectIdConvertError, err)
	}

	err = s.collection.FindOne(ctx, bson.M{"_id": obID}).Decode(&result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Students : ", result)
	return &result, nil
}

func (s *studentRepo) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	student.ID = primitive.NewObjectID()

	result, err := s.collection.InsertOne(ctx, student)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Students : ", result.InsertedID)
	student.ID = result.InsertedID.(primitive.ObjectID)
	return student, nil
}

func (s *studentRepo) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	_, err := s.collection.ReplaceOne(ctx, bson.M{"_id": student.ID}, student)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Students : ", student)
	return student, nil
}

func (s *studentRepo) DeleteStudent(ctx context.Context, id string) error {

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return err
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": obID})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Record Deleted Successfully", result.DeletedCount)
	return nil
}

func (s *studentRepo) SearchStudent(ctx context.Context, searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.StudentSearchData, error) {
	sortDirection := consts.SortASC
	if sortBy.Direction == consts.DESC {
		sortDirection = consts.SortDESC
	}
	var results []models.Student

	sortOptions := options.Find().SetSort(bson.M{sortBy.Column: sortDirection})
	paginationOptions := options.Find().SetSkip(int64(pagination.Page)).SetLimit(int64(pagination.PageSize))

	cursor, err := s.collection.Find(ctx,
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
	log.Info("Students : ", results)
	return nil, nil
}
