package student

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
	ctx          = context.Background()
	studentId, _ = primitive.ObjectIDFromHex("6454d4dd088ecfbcf32718c6")
	s1           = models.Student{
		ID:        studentId,
		FirstName: "test1",
		LastName:  "test1",
		Year:      1,
	}
	s2 = models.Student{
		ID:        studentId,
		FirstName: "test2",
		LastName:  "test2",
		Year:      2,
	}
	studentList = []models.Student{s1, s2}

	returnErr = errors.New("error")
)

func TestStudentUsecase_GetAllStudents_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected []models.Student
	}

	tests := []test{
		{
			expected: studentList,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindAllStudents(ctx).Return(studentList, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.GetAllStudents(ctx)
		if actual[0] != test.expected[0] || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}

}

func TestStudentUsecase_GetAllStudents_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindAllStudents(ctx).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.GetAllStudents(ctx)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_GetAllStudents(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindAllStudents(ctx).Return(studentList, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.GetAllStudents(ctx)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_GetStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindStudent(ctx, studentId).Return(&s1, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.GetStudent(ctx, studentId)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_GetStudent_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindStudent(ctx, studentId).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.GetStudent(ctx, studentId)
		if test.expected != err {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_GetStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().FindStudent(ctx, studentId).Return(&s1, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.GetStudent(ctx, studentId)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_CreateStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(ctx, &s1).Return(&s1, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.CreateStudent(ctx, &s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_CreateStudent_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(ctx, &s1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.CreateStudent(ctx, &s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_CreateStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(ctx, &s1).Return(&s1, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.CreateStudent(ctx, &s1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_UpdateStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s2,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(ctx, &s1).Return(&s2, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.UpdateStudent(ctx, &s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_UpdateStudent_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(ctx, &s1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.UpdateStudent(ctx, &s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_UpdateStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(ctx, &s1).Return(&s2, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.UpdateStudent(ctx, &s1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_DeleteStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: nil,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(ctx, studentId).Return(nil)

	student := NewStudent(repo)

	for _, test := range tests {
		err := student.DeleteStudent(ctx, studentId)
		if err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func TestStudentUsecase_DeleteStudent_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(ctx, studentId).Return(returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		err := student.DeleteStudent(ctx, studentId)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_DeleteStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(ctx, studentId).Return(nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		err := student.DeleteStudent(ctx, studentId)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_SearchStudents_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := models.StudentSearchData{
		TotalElements: 2,
		Data:          studentList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.StudentSearchData
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.SearchStudent(ctx, test.searchString, test.pagination, test.sortBy)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_SearchStudents_ErrorPath(t *testing.T) {
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.SearchStudent(ctx, test.searchString, test.pagination, test.sortBy)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_SearchStudents(b *testing.B) {
	ctrl := gomock.NewController(b)

	data := models.StudentSearchData{
		TotalElements: 2,
		Data:          studentList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.StudentSearchData
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

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(ctx, tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.SearchStudent(ctx, tests[0].searchString, tests[0].pagination,
			tests[0].sortBy)
		if err != nil {
			return
		}
	}
}
