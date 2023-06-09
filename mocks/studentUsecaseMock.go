// Code generated by MockGen. DO NOT EDIT.
// Source: ../src/github.com/shashaneRanasinghe/simpleMongoAPI/internal/usecases/student/studentUsecase.go

// Package mock_student is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/shashaneRanasinghe/simpleMongoAPI/internal/models"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockStudentUsecase is a mock of StudentUsecase interface.
type MockStudentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockStudentUsecaseMockRecorder
}

// MockStudentUsecaseMockRecorder is the mock recorder for MockStudentUsecase.
type MockStudentUsecaseMockRecorder struct {
	mock *MockStudentUsecase
}

// NewMockStudentUsecase creates a new mock instance.
func NewMockStudentUsecase(ctrl *gomock.Controller) *MockStudentUsecase {
	mock := &MockStudentUsecase{ctrl: ctrl}
	mock.recorder = &MockStudentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentUsecase) EXPECT() *MockStudentUsecaseMockRecorder {
	return m.recorder
}

// CreateStudent mocks base method.
func (m *MockStudentUsecase) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", ctx, student)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockStudentUsecaseMockRecorder) CreateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockStudentUsecase)(nil).CreateStudent), ctx, student)
}

// DeleteStudent mocks base method.
func (m *MockStudentUsecase) DeleteStudent(ctx context.Context, id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockStudentUsecaseMockRecorder) DeleteStudent(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockStudentUsecase)(nil).DeleteStudent), ctx, id)
}

// GetAllStudents mocks base method.
func (m *MockStudentUsecase) GetAllStudents(ctx context.Context) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStudents", ctx)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStudents indicates an expected call of GetAllStudents.
func (mr *MockStudentUsecaseMockRecorder) GetAllStudents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStudents", reflect.TypeOf((*MockStudentUsecase)(nil).GetAllStudents), ctx)
}

// GetStudent mocks base method.
func (m *MockStudentUsecase) GetStudent(ctx context.Context, id primitive.ObjectID) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudent", ctx, id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockStudentUsecaseMockRecorder) GetStudent(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockStudentUsecase)(nil).GetStudent), ctx, id)
}

// SearchStudent mocks base method.
func (m *MockStudentUsecase) SearchStudent(ctx context.Context, searchString string, pagination models.Pagination, sortBy models.SortBy) (*models.StudentSearchData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchStudent", ctx, searchString, pagination, sortBy)
	ret0, _ := ret[0].(*models.StudentSearchData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStudent indicates an expected call of SearchStudent.
func (mr *MockStudentUsecaseMockRecorder) SearchStudent(ctx, searchString, pagination, sortBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStudent", reflect.TypeOf((*MockStudentUsecase)(nil).SearchStudent), ctx, searchString, pagination, sortBy)
}

// UpdateStudent mocks base method.
func (m *MockStudentUsecase) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", ctx, student)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockStudentUsecaseMockRecorder) UpdateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockStudentUsecase)(nil).UpdateStudent), ctx, student)
}
