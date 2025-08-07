package service

import (
	"fmt"
	"schoolmgmt/services/student-service/internal/models"
	"schoolmgmt/services/student-service/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(req *models.CreateStudentRequest) (*models.Student, error) {
	// Check if student with email already exists
	_, err := s.repo.GetByEmail(req.Email)
	if err == nil {
		return nil, fmt.Errorf("student with email %s already exists", req.Email)
	}

	return s.repo.Create(req)
}

func (s *StudentService) GetStudent(id string) (*models.Student, error) {
	return s.repo.GetByID(id)
}

func (s *StudentService) UpdateStudent(id string, req *models.UpdateStudentRequest) (*models.Student, error) {
	return s.repo.Update(id, req)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.repo.Delete(id)
}

func (s *StudentService) ListStudents(limit, offset int) ([]*models.Student, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	
	return s.repo.List(limit, offset)
}
