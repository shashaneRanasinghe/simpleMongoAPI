package student

import (
	"context"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/repository"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/consts"
	"github.com/tryfix/log"
)

type StudentUsecase interface {
	GetAllStudents(ctx context.Context) ([]models.Student, error)
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	SearchStudent(ctx context.Context, searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.StudentSearchData, error)
	DeleteStudent(ctx context.Context, id string) error
}

type studentUsecase struct {
	studentRepo repository.StudentRepository
}

func NewStudent(studentRepo repository.StudentRepository) StudentUsecase {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}

func (s studentUsecase) GetAllStudents(ctx context.Context) ([]models.Student, error) {

	studentList, err := s.studentRepo.FindAllStudents(ctx)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return nil, err
	}
	return studentList, nil
}

func (s studentUsecase) GetStudent(ctx context.Context, id string) (*models.Student, error) {

	student, err := s.studentRepo.FindStudent(ctx, id)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return &models.Student{}, err
	}
	return student, nil
}

func (s studentUsecase) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	st, err := s.studentRepo.CreateStudent(ctx, student)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return &models.Student{}, err
	}
	return st, nil
}

func (s studentUsecase) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	st, err := s.studentRepo.UpdateStudent(ctx, student)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return &models.Student{}, err
	}
	return st, nil
}

func (s studentUsecase) SearchStudent(ctx context.Context, searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.StudentSearchData, error) {

	studentList, err := s.studentRepo.SearchStudent(ctx, searchString, pagination, sortBy)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return nil, err
	}
	return studentList, nil
}

func (s studentUsecase) DeleteStudent(ctx context.Context, id string) error {

	err := s.studentRepo.DeleteStudent(ctx, id)
	if err != nil {
		log.Debug(consts.StudentDeleteError, err)
		return err
	}
	return nil
}
