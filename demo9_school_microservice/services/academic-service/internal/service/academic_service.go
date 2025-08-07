package service

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"schoolmgmt/services/academic-service/internal/models"
	"schoolmgmt/services/academic-service/internal/repository"
)

type AcademicService struct {
	repo   *repository.AcademicRepository
	logger *logrus.Logger
}

func NewAcademicService(repo *repository.AcademicRepository) *AcademicService {
	return &AcademicService{
		repo:   repo,
		logger: logrus.New(),
	}
}

func (s *AcademicService) CreateAcademic(req *models.CreateAcademicRequest) (*repository.Academic, error) {
	s.logger.WithFields(logrus.Fields{
		"operation":  "service_create_academic",
		"student_id": req.StudentID,
		"teacher_id": req.TeacherID,
		"subject":    req.Subject,
	}).Info("Creating academic record via service")

	academic := &repository.Academic{
		StudentID:     req.StudentID,
		TeacherID:     req.TeacherID,
		Subject:       req.Subject,
		Grade:         req.Grade,
		Semester:      req.Semester,
		AcademicYear:  req.AcademicYear,
		ExamType:      req.ExamType,
		MaxMarks:      req.MaxMarks,
		ObtainedMarks: req.ObtainedMarks,
		Remarks:       req.Remarks,
		ExamDate:      time.Now(),
	}

	// Calculate percentage
	if academic.MaxMarks > 0 {
		academic.Percentage = (academic.ObtainedMarks / academic.MaxMarks) * 100
	}

	// Determine status
	if academic.Percentage >= 60 {
		academic.Status = "pass"
	} else {
		academic.Status = "fail"
	}

	err := s.repo.CreateAcademic(academic)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create academic record")
		return nil, err
	}

	s.logger.WithField("academic_id", academic.ID).Info("Academic record created successfully")
	return academic, nil
}

func (s *AcademicService) GetAcademic(id string) (*repository.Academic, error) {
	s.logger.WithField("academic_id", id).Info("Retrieving academic record")
	
	academic, err := s.repo.GetAcademicByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("academic_id", id).Error("Failed to retrieve academic record")
		return nil, err
	}

	return academic, nil
}

func (s *AcademicService) UpdateAcademic(id string, req *models.UpdateAcademicRequest) (*repository.Academic, error) {
	s.logger.WithField("academic_id", id).Info("Updating academic record")

	academic, err := s.repo.GetAcademicByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Grade != nil {
		academic.Grade = *req.Grade
	}
	if req.ObtainedMarks != nil {
		academic.ObtainedMarks = *req.ObtainedMarks
		// Recalculate percentage
		if academic.MaxMarks > 0 {
			academic.Percentage = (*req.ObtainedMarks / academic.MaxMarks) * 100
		}
		// Update status
		if academic.Percentage >= 60 {
			academic.Status = "pass"
		} else {
			academic.Status = "fail"
		}
	}
	if req.Remarks != nil {
		academic.Remarks = *req.Remarks
	}

	academic.UpdatedAt = time.Now()

	err = s.repo.UpdateAcademic(academic)
	if err != nil {
		s.logger.WithError(err).Error("Failed to update academic record")
		return nil, err
	}

	return academic, nil
}

func (s *AcademicService) DeleteAcademic(id string) error {
	s.logger.WithField("academic_id", id).Info("Deleting academic record")
	
	err := s.repo.DeleteAcademic(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to delete academic record")
		return err
	}

	return nil
}

func (s *AcademicService) ListAcademics(limit, offset int) ([]*repository.Academic, error) {
	s.logger.WithFields(logrus.Fields{
		"limit":  limit,
		"offset": offset,
	}).Info("Listing academic records")

	academics, err := s.repo.ListAcademics(limit, offset)
	if err != nil {
		s.logger.WithError(err).Error("Failed to list academic records")
		return nil, err
	}

	s.logger.WithField("count", len(academics)).Info("Academic records retrieved successfully")
	return academics, nil
}

func (s *AcademicService) CreateClass(req *models.CreateClassRequest) (*repository.Class, error) {
	s.logger.WithFields(logrus.Fields{
		"operation":   "service_create_class",
		"class_name":  req.ClassName,
		"teacher_id":  req.TeacherID,
	}).Info("Creating class via service")

	class := &repository.Class{
		ClassName:    req.ClassName,
		Grade:        req.Grade,
		Section:      req.Section,
		TeacherID:    req.TeacherID,
		Subject:      req.Subject,
		AcademicYear: req.AcademicYear,
		Semester:     req.Semester,
		StudentIDs:   req.StudentIDs,
		MaxCapacity:  req.MaxCapacity,
		Status:       "active",
	}

	err := s.repo.CreateClass(class)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create class")
		return nil, err
	}

	s.logger.WithField("class_id", class.ID).Info("Class created successfully")
	return class, nil
}

func (s *AcademicService) GetClass(id string) (*repository.Class, error) {
	s.logger.WithField("class_id", id).Info("Retrieving class")
	
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("class_id", id).Error("Failed to retrieve class")
		return nil, err
	}

	return class, nil
}

func (s *AcademicService) ListClasses(limit, offset int) ([]*repository.Class, error) {
	s.logger.WithFields(logrus.Fields{
		"limit":  limit,
		"offset": offset,
	}).Info("Listing classes")

	classes, err := s.repo.ListClasses(limit, offset)
	if err != nil {
		s.logger.WithError(err).Error("Failed to list classes")
		return nil, err
	}

	s.logger.WithField("count", len(classes)).Info("Classes retrieved successfully")
	return classes, nil
}
