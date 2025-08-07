package service

import (
	"context"
	"log"

	"achievement-service/internal/models"
	"achievement-service/internal/repository"
)

type AchievementService struct {
	repo repository.AchievementRepository
}

func NewAchievementService(repo repository.AchievementRepository) *AchievementService {
	log.Println("AchievementService: Creating new achievement service with Couchbase repository")
	return &AchievementService{
		repo: repo,
	}
}

func (s *AchievementService) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
	log.Printf("AchievementService: Creating achievement - ID: %s, Type: %s, Student: %s", 
		achievement.ID, achievement.Type, achievement.StudentID)
	
	err := s.repo.Create(ctx, achievement)
	if err != nil {
		log.Printf("AchievementService: Error creating achievement %s: %v", achievement.ID, err)
		return err
	}
	
	log.Printf("AchievementService: Successfully created achievement %s in Couchbase", achievement.ID)
	return nil
}

func (s *AchievementService) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
	log.Printf("AchievementService: Retrieving achievement with ID: %s", id)
	
	achievement, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("AchievementService: Error retrieving achievement %s: %v", id, err)
		return nil, err
	}
	
	log.Printf("AchievementService: Successfully retrieved achievement %s from Couchbase", id)
	return achievement, nil
}

func (s *AchievementService) GetAchievementsByStudent(ctx context.Context, studentID string) ([]*models.Achievement, error) {
	log.Printf("AchievementService: Retrieving achievements for student: %s", studentID)
	
	achievements, err := s.repo.GetByStudentID(ctx, studentID)
	if err != nil {
		log.Printf("AchievementService: Error retrieving achievements for student %s: %v", studentID, err)
		return nil, err
	}
	
	log.Printf("AchievementService: Successfully retrieved %d achievements for student %s from Couchbase", 
		len(achievements), studentID)
	return achievements, nil
}

func (s *AchievementService) UpdateAchievement(ctx context.Context, achievement *models.Achievement) error {
	log.Printf("AchievementService: Updating achievement - ID: %s, Type: %s", 
		achievement.ID, achievement.Type)
	
	err := s.repo.Update(ctx, achievement)
	if err != nil {
		log.Printf("AchievementService: Error updating achievement %s: %v", achievement.ID, err)
		return err
	}
	
	log.Printf("AchievementService: Successfully updated achievement %s in Couchbase", achievement.ID)
	return nil
}

func (s *AchievementService) DeleteAchievement(ctx context.Context, id string) error {
	log.Printf("AchievementService: Deleting achievement with ID: %s", id)
	
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Printf("AchievementService: Error deleting achievement %s: %v", id, err)
		return err
	}
	
	log.Printf("AchievementService: Successfully deleted achievement %s from Couchbase", id)
	return nil
}

func (s *AchievementService) GetAllAchievements(ctx context.Context) ([]*models.Achievement, error) {
	log.Println("AchievementService: Retrieving all achievements")
	
	achievements, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("AchievementService: Error retrieving all achievements: %v", err)
		return nil, err
	}
	
	log.Printf("AchievementService: Successfully retrieved %d achievements from Couchbase", len(achievements))
	return achievements, nil
}
