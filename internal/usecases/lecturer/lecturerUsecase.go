package lecturer

import (
	"context"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/repository"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/consts"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LecturerUsecase interface {
	GetAllLecturers(ctx context.Context) ([]models.Lecturer, error)
	GetLecturer(ctx context.Context, id primitive.ObjectID) (*models.Lecturer, error)
	CreateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error)
	UpdateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error)
	SearchLecturer(ctx context.Context, searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.LecturerSearchData, error)
	DeleteLecturer(ctx context.Context, id primitive.ObjectID) error
}

type lecturerUsecase struct {
	lecturerRepo repository.LecturerRepository
}

func NewLecturer(lecturerRepo repository.LecturerRepository) LecturerUsecase {
	return &lecturerUsecase{
		lecturerRepo: lecturerRepo,
	}
}

func (s lecturerUsecase) GetAllLecturers(ctx context.Context) ([]models.Lecturer, error) {

	lecturerList, err := s.lecturerRepo.FindAllLecturers(ctx)
	if err != nil {
		log.Debug(consts.LecturerGetError, err)
		return nil, err
	}
	return lecturerList, nil
}

func (s lecturerUsecase) GetLecturer(ctx context.Context, id primitive.ObjectID) (*models.Lecturer, error) {

	lecturer, err := s.lecturerRepo.FindLecturer(ctx, id)
	if err != nil {
		log.Debug(consts.LecturerGetError, err)
		return &models.Lecturer{}, err
	}
	return lecturer, nil
}

func (s lecturerUsecase) CreateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error) {

	st, err := s.lecturerRepo.CreateLecturer(ctx, lecturer)
	if err != nil {
		log.Debug(consts.LecturerCreateError, err)
		return &models.Lecturer{}, err
	}
	return st, nil
}

func (s lecturerUsecase) UpdateLecturer(ctx context.Context, lecturer *models.Lecturer) (*models.Lecturer, error) {

	st, err := s.lecturerRepo.UpdateLecturer(ctx, lecturer)
	if err != nil {
		log.Debug(consts.LecturerUpdateError, err)
		return &models.Lecturer{}, err
	}
	return st, nil
}

func (s lecturerUsecase) SearchLecturer(ctx context.Context, searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.LecturerSearchData, error) {

	lecturerList, err := s.lecturerRepo.SearchLecturer(ctx, searchString, pagination, sortBy)
	if err != nil {
		log.Debug(consts.LecturerGetError, err)
		return nil, err
	}
	return lecturerList, nil
}

func (s lecturerUsecase) DeleteLecturer(ctx context.Context, id primitive.ObjectID) error {

	err := s.lecturerRepo.DeleteLecturer(ctx, id)
	if err != nil {
		log.Debug(consts.LecturerDeleteError, err)
		return err
	}
	return nil
}
