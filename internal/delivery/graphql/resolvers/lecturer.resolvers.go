package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	models2 "github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/consts"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/delivery/graphql/models"
)

// CreateLecturer is the resolver for the createLecturer field.
func (r *mutationResolver) CreateLecturer(ctx context.Context, lecturer *models.LecturerInput) (*models.LecturerResponse, error) {
	var resp models.LecturerResponse

	st := models2.Lecturer{
		FirstName: *lecturer.Firstname,
		LastName:  *lecturer.Lastname,
		Year:      *lecturer.Year,
	}

	stResult, err := r.lecturerUsecase.CreateLecturer(ctx, &st)
	if err != nil {
		log.Error(consts.LecturerGetError, err)
		resp.Status = consts.Error
		resp.Data = nil
		resp.Message = consts.LecturerGetError
		return &resp, err
	}

	respLecturer := models.Lecturer{
		ID:        stResult.ID.Hex(),
		Firstname: stResult.FirstName,
		Lastname:  stResult.LastName,
		Year:      stResult.Year,
	}

	resp.Status = consts.Success
	resp.Data = &respLecturer
	resp.Message = consts.LecturerGetSuccess
	return &resp, err
}

// UpdateLecturer is the resolver for the updateLecturer field.
func (r *mutationResolver) UpdateLecturer(ctx context.Context, lecturer *models.LecturerInput) (*models.LecturerResponse, error) {
	var resp models.LecturerResponse
	id, err := primitive.ObjectIDFromHex(*lecturer.ID)
	if err != nil {
		log.Error(consts.IDError, err)
		resp.Status = consts.Error
		resp.Data = nil
		resp.Message = consts.IDError
		return &resp, err
	}
	st := models2.Lecturer{
		ID:        id,
		FirstName: *lecturer.Firstname,
		LastName:  *lecturer.Lastname,
		Year:      *lecturer.Year,
	}

	stResult, err := r.lecturerUsecase.UpdateLecturer(ctx, &st)
	if err != nil {
		log.Error(consts.LecturerUpdateError, err)
		resp.Status = consts.Error
		resp.Data = nil
		resp.Message = consts.LecturerUpdateError
		return &resp, err
	}

	respLecturer := models.Lecturer{
		ID:        stResult.ID.Hex(),
		Firstname: stResult.FirstName,
		Lastname:  stResult.LastName,
		Year:      stResult.Year,
	}

	resp.Status = consts.Success
	resp.Data = &respLecturer
	resp.Message = consts.LecturerUpdateSuccess
	return &resp, err
}

// DeleteLecturer is the resolver for the deleteLecturer field.
func (r *mutationResolver) DeleteLecturer(ctx context.Context, lecturerID string) (*models.LecturerDeleteResponse, error) {
	var resp models.LecturerDeleteResponse
	id, err := primitive.ObjectIDFromHex(lecturerID)
	if err != nil {
		log.Error(consts.ObjectIdConvertError, err)
		resp.Status = consts.Error
		resp.Message = consts.ObjectIdConvertError
		return &resp, err
	}
	err = r.lecturerUsecase.DeleteLecturer(ctx, id)
	if err != nil {
		log.Error(consts.LecturerDeleteError, err)
		resp.Status = consts.Error
		resp.Message = consts.LecturerDeleteError
		return &resp, err
	}
	resp.Status = consts.Success
	resp.Message = consts.LecturerDeleteSuccess
	return &resp, err
}

// GetAllLecturers is the resolver for the getAllLecturers field.
func (r *queryResolver) GetAllLecturers(ctx context.Context) (*models.LecturerListResponse, error) {
	var resp models.LecturerListResponse
	var resLecturers []*models.Lecturer

	lecturers, err := r.lecturerUsecase.GetAllLecturers(ctx)
	if err != nil {
		log.Error(consts.LecturerGetError, err)
		resp.Status = consts.Error
		resp.Data = nil
		resp.Message = consts.LecturerGetError
		return &resp, err
	}

	for _, st := range lecturers {
		resLecturer := models.Lecturer{
			ID:        (st.ID).Hex(),
			Firstname: st.FirstName,
			Lastname:  st.LastName,
			Year:      st.Year,
		}
		resLecturers = append(resLecturers, &resLecturer)
	}

	resp.Status = consts.Success
	resp.Data = resLecturers
	resp.Message = consts.LecturerGetSuccess
	return &resp, err
}

// GetLecturer is the resolver for the getLecturer field.
func (r *queryResolver) GetLecturer(ctx context.Context, lecturerID string) (*models.LecturerResponse, error) {
	var resp models.LecturerResponse
	id, err := primitive.ObjectIDFromHex(lecturerID)
	if err != nil {
		log.Error(consts.ObjectIdConvertError, err)
		resp.Status = consts.Error
		resp.Message = consts.ObjectIdConvertError
		return &resp, err
	}
	reLecturer, err := r.lecturerUsecase.GetLecturer(ctx, id)
	if err != nil {
		log.Error(consts.LecturerGetError, err)
		resp.Status = consts.Error
		resp.Message = consts.LecturerGetError
		return &resp, err
	}
	st := models.Lecturer{
		ID:        reLecturer.ID.Hex(),
		Firstname: reLecturer.FirstName,
		Lastname:  reLecturer.LastName,
		Year:      reLecturer.Year,
	}
	resp.Status = consts.Success
	resp.Data = &st
	resp.Message = consts.LecturerGetSuccess
	return &resp, err
}

// SearchLecturer is the resolver for the searchLecturer field.
func (r *queryResolver) SearchLecturer(ctx context.Context, searchString *string, pagination *models.Pagination, sortBy *models.SortBy) (*models.LecturerSearchResponse, error) {
	var resp models.LecturerSearchResponse
	var lecturerSearch models.LecturerSearch
	var resLecturers []*models.Lecturer

	paginationIn := models2.Pagination{
		Page:     *pagination.Page,
		PageSize: *pagination.PageSize,
	}
	sortByIn := models2.SortBy{
		Column:    *sortBy.Column,
		Direction: string(*sortBy.Direction),
	}

	searchResult, err := r.lecturerUsecase.SearchLecturer(ctx, *searchString, paginationIn, sortByIn)
	if err != nil {
		log.Error(consts.LecturerGetError, err)
		resp.Status = consts.Error
		resp.Message = consts.LecturerGetError
		return &resp, err
	}
	for _, st := range searchResult.Data {
		resLecturer := models.Lecturer{
			ID:        (st.ID).Hex(),
			Firstname: st.FirstName,
			Lastname:  st.LastName,
			Year:      st.Year,
		}
		resLecturers = append(resLecturers, &resLecturer)
	}

	lecturerSearch.TotalElements = searchResult.TotalElements
	lecturerSearch.Data = resLecturers

	resp.Status = consts.Success
	resp.Data = &lecturerSearch
	resp.Message = consts.LecturerGetSuccess
	return &resp, err
}