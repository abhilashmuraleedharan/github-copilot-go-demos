// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package service

import (
	"context"
	"errors"
	"school-microservice/internal/models"
	"school-microservice/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStudentRepository is a mock implementation of StudentRepository
type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) Create(ctx context.Context, student *models.Student) error {
	args := m.Called(ctx, student)
	return args.Error(0)
}

func (m *MockStudentRepository) GetByID(ctx context.Context, id string) (*models.Student, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentRepository) GetAll(ctx context.Context, page, pageSize int) ([]*models.Student, int, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Int(1), args.Error(2)
	}
	return args.Get(0).([]*models.Student), args.Int(1), args.Error(2)
}

func (m *MockStudentRepository) Update(ctx context.Context, id string, student *models.Student) error {
	args := m.Called(ctx, id, student)
	return args.Error(0)
}

func (m *MockStudentRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStudentRepository) GetByGrade(ctx context.Context, grade string) ([]*models.Student, error) {
	args := m.Called(ctx, grade)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Student), args.Error(1)
}

func (m *MockStudentRepository) GetByEmail(ctx context.Context, email string) (*models.Student, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Student), args.Error(1)
}

// MockRepository is a mock implementation of Repository
type MockRepository struct {
	Student *MockStudentRepository
}

// Helper function to create a test service with mocked repository
func setupStudentService() (*studentService, *MockStudentRepository) {
	mockStudentRepo := &MockStudentRepository{}
	mockRepo := &repository.Repository{
		Student: mockStudentRepo,
	}
	service := &studentService{repo: mockRepo}
	return service, mockStudentRepo
}

// Helper function to create a valid student request
func createValidStudentRequest() *models.CreateStudentRequest {
	dateOfBirth := time.Date(2010, time.May, 15, 0, 0, 0, 0, time.UTC)
	return &models.CreateStudentRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: dateOfBirth,
		Grade:       "8",
		Address:     "123 Main St",
		Phone:       "1234567890",
		ParentName:  "Jane Doe",
		ParentPhone: "0987654321",
	}
}

// Helper function to create a student model
func createStudentModel() *models.Student {
	dateOfBirth := time.Date(2010, time.May, 15, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	return &models.Student{
		ID:          "student_123",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: dateOfBirth,
		Grade:       "8",
		Address:     "123 Main St",
		Phone:       "1234567890",
		ParentName:  "Jane Doe",
		ParentPhone: "0987654321",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func TestStudentService_CreateStudent_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()

	// Setup mock expectations
	mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, errors.New("not found"))
	mockRepo.On("Create", ctx, mock.AnythingOfType("*models.Student")).Return(nil)

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.FirstName, result.FirstName)
	assert.Equal(t, req.LastName, result.LastName)
	assert.Equal(t, req.Email, result.Email)
	assert.Equal(t, req.Grade, result.Grade)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_CreateStudent_ValidationError_MissingFirstName(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.FirstName = ""

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "first name is required")
}

func TestStudentService_CreateStudent_ValidationError_MissingLastName(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.LastName = ""

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "last name is required")
}

func TestStudentService_CreateStudent_ValidationError_InvalidEmail(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.Email = "invalid-email"

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid email format")
}

func TestStudentService_CreateStudent_ValidationError_InvalidGrade(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.Grade = "invalid-grade"

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid grade")
}

func TestStudentService_CreateStudent_ValidationError_FutureDateOfBirth(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.DateOfBirth = time.Now().Add(24 * time.Hour) // Tomorrow

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "date of birth cannot be in the future")
}

func TestStudentService_CreateStudent_ValidationError_InvalidAge(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	req.DateOfBirth = time.Now().AddDate(-30, 0, 0) // 30 years ago

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student age must be between 3 and 25 years")
}

func TestStudentService_CreateStudent_DuplicateEmail(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()
	existingStudent := createStudentModel()

	// Setup mock expectations
	mockRepo.On("GetByEmail", ctx, req.Email).Return(existingStudent, nil)

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student with email")
	assert.Contains(t, err.Error(), "already exists")
	mockRepo.AssertExpectations(t)
}

func TestStudentService_CreateStudent_RepositoryError(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()

	// Setup mock expectations
	mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, errors.New("not found"))
	mockRepo.On("Create", ctx, mock.AnythingOfType("*models.Student")).Return(errors.New("database error"))

	// Act
	result, err := service.CreateStudent(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create student")
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetStudent_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "student_123"
	expectedStudent := createStudentModel()

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(expectedStudent, nil)

	// Act
	result, err := service.GetStudent(ctx, studentID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedStudent.ID, result.ID)
	assert.Equal(t, expectedStudent.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetStudent_EmptyID(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()

	// Act
	result, err := service.GetStudent(ctx, "")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student ID is required")
}

func TestStudentService_GetStudent_NotFound(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "nonexistent"

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(nil, errors.New("not found"))

	// Act
	result, err := service.GetStudent(ctx, studentID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetAllStudents_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	page, pageSize := 1, 10
	students := []*models.Student{createStudentModel()}
	totalCount := 1

	// Setup mock expectations
	mockRepo.On("GetAll", ctx, page, pageSize).Return(students, totalCount, nil)

	// Act
	result, total, err := service.GetAllStudents(ctx, page, pageSize)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, totalCount, total)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetAllStudents_DefaultPagination(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	students := []*models.Student{createStudentModel()}
	totalCount := 1

	// Setup mock expectations - should default to page 1, pageSize 10
	mockRepo.On("GetAll", ctx, 1, 10).Return(students, totalCount, nil)

	// Act
	result, total, err := service.GetAllStudents(ctx, 0, 0) // Invalid values should be defaulted

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, totalCount, total)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetAllStudents_MaxPageSize(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	students := []*models.Student{createStudentModel()}
	totalCount := 1

	// Setup mock expectations - should limit pageSize to 10 (default)
	mockRepo.On("GetAll", ctx, 1, 10).Return(students, totalCount, nil)

	// Act
	result, total, err := service.GetAllStudents(ctx, 1, 500) // Large pageSize should be limited

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, totalCount, total)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_UpdateStudent_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "student_123"
	req := createValidStudentRequest()
	existingStudent := createStudentModel()

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(existingStudent, nil)
	mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, errors.New("not found")) // Email not found (can update)
	mockRepo.On("Update", ctx, studentID, mock.AnythingOfType("*models.Student")).Return(nil)

	// Act
	result, err := service.UpdateStudent(ctx, studentID, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, studentID, result.ID)
	assert.Equal(t, req.FirstName, result.FirstName)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_UpdateStudent_EmptyID(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()
	req := createValidStudentRequest()

	// Act
	result, err := service.UpdateStudent(ctx, "", req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student ID is required")
}

func TestStudentService_UpdateStudent_StudentNotFound(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "nonexistent"
	req := createValidStudentRequest()

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(nil, errors.New("not found"))

	// Act
	result, err := service.UpdateStudent(ctx, studentID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student not found")
	mockRepo.AssertExpectations(t)
}

func TestStudentService_UpdateStudent_EmailAlreadyExists(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "student_123"
	req := createValidStudentRequest()
	req.Email = "different@example.com"
	existingStudent := createStudentModel()
	existingStudentWithSameEmail := createStudentModel()
	existingStudentWithSameEmail.ID = "different_student"
	existingStudentWithSameEmail.Email = req.Email

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(existingStudent, nil)
	mockRepo.On("GetByEmail", ctx, req.Email).Return(existingStudentWithSameEmail, nil)

	// Act
	result, err := service.UpdateStudent(ctx, studentID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student with email")
	assert.Contains(t, err.Error(), "already exists")
	mockRepo.AssertExpectations(t)
}

func TestStudentService_DeleteStudent_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "student_123"
	existingStudent := createStudentModel()

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(existingStudent, nil)
	mockRepo.On("Delete", ctx, studentID).Return(nil)

	// Act
	err := service.DeleteStudent(ctx, studentID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_DeleteStudent_EmptyID(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()

	// Act
	err := service.DeleteStudent(ctx, "")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "student ID is required")
}

func TestStudentService_DeleteStudent_StudentNotFound(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	studentID := "nonexistent"

	// Setup mock expectations
	mockRepo.On("GetByID", ctx, studentID).Return(nil, errors.New("not found"))

	// Act
	err := service.DeleteStudent(ctx, studentID)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "student not found")
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetStudentsByGrade_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	grade := "8"
	students := []*models.Student{createStudentModel()}

	// Setup mock expectations
	mockRepo.On("GetByGrade", ctx, grade).Return(students, nil)

	// Act
	result, err := service.GetStudentsByGrade(ctx, grade)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, grade, result[0].Grade)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetStudentsByGrade_EmptyGrade(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()

	// Act
	result, err := service.GetStudentsByGrade(ctx, "")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "grade is required")
}

func TestStudentService_GetStudentByEmail_Success(t *testing.T) {
	// Arrange
	service, mockRepo := setupStudentService()
	ctx := context.Background()
	email := "john.doe@example.com"
	expectedStudent := createStudentModel()

	// Setup mock expectations
	mockRepo.On("GetByEmail", ctx, email).Return(expectedStudent, nil)

	// Act
	result, err := service.GetStudentByEmail(ctx, email)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestStudentService_GetStudentByEmail_EmptyEmail(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	ctx := context.Background()

	// Act
	result, err := service.GetStudentByEmail(ctx, "")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "email is required")
}

// Test validation helper function
func TestStudentService_validateStudentRequest_AllValidationErrors(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	req := &models.CreateStudentRequest{
		FirstName:   "",                                     // Invalid: empty
		LastName:    "",                                     // Invalid: empty
		Email:       "invalid-email",                        // Invalid: format
		DateOfBirth: time.Now().Add(24 * time.Hour),        // Invalid: future date
		Grade:       "invalid",                              // Invalid: not in valid range
		Phone:       "123",                                  // Invalid: too short
		ParentPhone: "456",                                  // Invalid: too short
	}

	// Act
	err := service.validateStudentRequest(req)

	// Assert
	assert.Error(t, err)
	// The function returns the first validation error, so we just check that it's a validation error
	assert.True(t, len(err.Error()) > 0)
}

func TestStudentService_validateStudentRequest_ValidGrades(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	validGrades := []string{"KG", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}

	for _, grade := range validGrades {
		t.Run("Grade_"+grade, func(t *testing.T) {
			req := createValidStudentRequest()
			req.Grade = grade

			// Act
			err := service.validateStudentRequest(req)

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestStudentService_validateStudentRequest_OptionalFields(t *testing.T) {
	// Arrange
	service, _ := setupStudentService()
	req := createValidStudentRequest()
	req.Address = ""      // Optional field
	req.Phone = ""        // Optional field
	req.ParentName = ""   // Optional field
	req.ParentPhone = ""  // Optional field

	// Act
	err := service.validateStudentRequest(req)

	// Assert
	assert.NoError(t, err)
}