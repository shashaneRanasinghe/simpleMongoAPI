package lecturer

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleMongoAPI/mocks"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var (
	ctx           = context.Background()
	lecturerId, _ = primitive.ObjectIDFromHex("6454d4dd088ecfbcf32718c6")
	s1            = models.Lecturer{
		ID:        lecturerId,
		FirstName: "test1",
		LastName:  "test1",
		Year:      1,
	}
	s2 = models.Lecturer{
		ID:        lecturerId,
		FirstName: "test2",
		LastName:  "test2",
		Year:      2,
	}
	lecturerList = []models.Lecturer{s1, s2}

	returnErr = errors.New("error")
)

func TestLecturerUsecase_GetAllLecturers_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected []models.Lecturer
	}

	tests := []test{
		{
			expected: lecturerList,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindAllLecturers(ctx).Return(lecturerList, nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.GetAllLecturers(ctx)
		if actual[0] != test.expected[0] || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}

}

func TestLecturerUsecase_GetAllLecturers_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindAllLecturers(ctx).Return(nil, returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		_, err := lecturer.GetAllLecturers(ctx)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_GetAllLecturers(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindAllLecturers(ctx).Return(lecturerList, nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		_, err := lecturer.GetAllLecturers(ctx)
		if err != nil {
			return
		}
	}
}

func TestLecturerUsecase_GetLecturer_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Lecturer
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindLecturer(ctx, lecturerId).Return(&s1, nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.GetLecturer(ctx, lecturerId)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestLecturerUsecase_GetLecturer_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindLecturer(ctx, lecturerId).Return(nil, returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		_, err := lecturer.GetLecturer(ctx, lecturerId)
		if test.expected != err {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_GetLecturer(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().FindLecturer(ctx, lecturerId).Return(&s1, nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		_, err := lecturer.GetLecturer(ctx, lecturerId)
		if err != nil {
			return
		}
	}
}

func TestLecturerUsecase_CreateLecturer_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Lecturer
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().CreateLecturer(ctx, &s1).Return(&s1, nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.CreateLecturer(ctx, &s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestLecturerUsecase_CreateLecturer_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().CreateLecturer(ctx, &s1).Return(nil, returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		_, err := lecturer.CreateLecturer(ctx, &s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_CreateLecturer(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().CreateLecturer(ctx, &s1).Return(&s1, nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		_, err := lecturer.CreateLecturer(ctx, &s1)
		if err != nil {
			return
		}
	}
}

func TestLecturerUsecase_UpdateLecturer_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Lecturer
	}

	tests := []test{
		{
			expected: &s2,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().UpdateLecturer(ctx, &s1).Return(&s2, nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.UpdateLecturer(ctx, &s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestLecturerUsecase_UpdateLecturer_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().UpdateLecturer(ctx, &s1).Return(nil, returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		_, err := lecturer.UpdateLecturer(ctx, &s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_UpdateLecturer(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().UpdateLecturer(ctx, &s1).Return(&s2, nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		_, err := lecturer.UpdateLecturer(ctx, &s1)
		if err != nil {
			return
		}
	}
}

func TestLecturerUsecase_DeleteLecturer_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Lecturer
	}

	tests := []test{
		{
			expected: nil,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().DeleteLecturer(ctx, lecturerId).Return(nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		err := lecturer.DeleteLecturer(ctx, lecturerId)
		if err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func TestLecturerUsecase_DeleteLecturer_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().DeleteLecturer(ctx, lecturerId).Return(returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		err := lecturer.DeleteLecturer(ctx, lecturerId)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_DeleteLecturer(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().DeleteLecturer(ctx, lecturerId).Return(nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		err := lecturer.DeleteLecturer(ctx, lecturerId)
		if err != nil {
			return
		}
	}
}

func TestLecturerUsecase_SearchLecturers_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := models.LecturerSearchData{
		TotalElements: 2,
		Data:          lecturerList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.LecturerSearchData
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: &data,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().SearchLecturer(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.SearchLecturer(ctx, test.searchString, test.pagination, test.sortBy)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestLecturerUsecase_SearchLecturers_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     error
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: returnErr,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().SearchLecturer(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(nil, returnErr)

	lecturer := NewLecturer(repo)

	for _, test := range tests {
		actual, err := lecturer.SearchLecturer(ctx, test.searchString, test.pagination, test.sortBy)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkLecturerUsecase_SearchLecturers(b *testing.B) {
	ctrl := gomock.NewController(b)

	data := models.LecturerSearchData{
		TotalElements: 2,
		Data:          lecturerList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.LecturerSearchData
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: &data,
		},
	}

	repo := mocks.NewMockLecturerRepository(ctrl)
	repo.EXPECT().SearchLecturer(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil).AnyTimes()

	lecturer := NewLecturer(repo)

	for i := 0; i < b.N; i++ {
		_, err := lecturer.SearchLecturer(ctx, tests[0].searchString, tests[0].pagination,
			tests[0].sortBy)
		if err != nil {
			return
		}
	}
}
