package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
)

type Achievement struct {
	ID          string    `json:"id"`
	StudentID   string    `json:"student_id"`
	TeacherID   string    `json:"teacher_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // academic, sports, cultural, leadership, community
	Type        string    `json:"achievement_type"` // award, certificate, medal, recognition (renamed to avoid conflict with document type)
	Level       string    `json:"level"`    // school, district, state, national, international
	Date        time.Time `json:"date"`
	Points      int       `json:"points"`
	Status      string    `json:"status"` // pending, approved, rejected
	Remarks     string    `json:"remarks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DocType     string    `json:"type"` // For document type identification in Couchbase
}

type Award struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Criteria    string    `json:"criteria"`
	Points      int       `json:"points"`
	Level       string    `json:"level"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"` // For document type identification
}

type AchievementRepository struct {
	collection *gocb.Collection
	cluster    *gocb.Cluster
	logger     *logrus.Logger
}

func NewAchievementRepository(cluster *gocb.Cluster, collection *gocb.Collection, logger *logrus.Logger) *AchievementRepository {
	return &AchievementRepository{
		collection: collection,
		cluster:    cluster,
		logger:     logger,
	}
}

// Achievement CRUD operations
func (r *AchievementRepository) CreateAchievement(achievement *Achievement) error {
	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_create",
		"achievement_id": achievement.ID,
		"student_id":     achievement.StudentID,
		"title":          achievement.Title,
	}).Info("Creating new achievement in Couchbase")

	// Set metadata
	achievement.CreatedAt = time.Now()
	achievement.UpdatedAt = time.Now()
	achievement.DocType = "achievement"

	// Generate ID if not provided
	if achievement.ID == "" {
		achievement.ID = fmt.Sprintf("achievement_%d", time.Now().UnixNano())
		r.logger.WithField("generated_id", achievement.ID).Info("Generated new achievement ID")
	}

	// Set default status if not provided
	if achievement.Status == "" {
		achievement.Status = "pending"
	}

	documentKey := fmt.Sprintf("achievement::%s", achievement.ID)
	
	_, err := r.collection.Insert(documentKey, achievement, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":      "achievement_create",
			"achievement_id": achievement.ID,
			"document_key":   documentKey,
			"error":          err.Error(),
		}).Error("Failed to create achievement in Couchbase")
		return fmt.Errorf("failed to create achievement: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_create",
		"achievement_id": achievement.ID,
		"document_key":   documentKey,
	}).Info("Successfully created achievement in Couchbase")

	return nil
}

func (r *AchievementRepository) GetAchievementByID(id string) (*Achievement, error) {
	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_get_by_id",
		"achievement_id": id,
	}).Info("Retrieving achievement from Couchbase")

	documentKey := fmt.Sprintf("achievement::%s", id)
	
	result, err := r.collection.Get(documentKey, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":      "achievement_get_by_id",
			"achievement_id": id,
			"document_key":   documentKey,
			"error":          err.Error(),
		}).Error("Failed to retrieve achievement from Couchbase")
		return nil, fmt.Errorf("failed to get achievement: %w", err)
	}

	var achievement Achievement
	if err := result.Content(&achievement); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":      "achievement_get_by_id",
			"achievement_id": id,
			"document_key":   documentKey,
			"error":          err.Error(),
		}).Error("Failed to decode achievement document")
		return nil, fmt.Errorf("failed to decode achievement: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_get_by_id",
		"achievement_id": id,
		"document_key":   documentKey,
	}).Info("Successfully retrieved achievement from Couchbase")

	return &achievement, nil
}

func (r *AchievementRepository) UpdateAchievement(id string, achievement *Achievement) error {
	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_update",
		"achievement_id": id,
	}).Info("Updating achievement in Couchbase")

	// Ensure the ID matches
	achievement.ID = id
	achievement.UpdatedAt = time.Now()
	achievement.DocType = "achievement"

	documentKey := fmt.Sprintf("achievement::%s", id)
	
	_, err := r.collection.Replace(documentKey, achievement, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":      "achievement_update",
			"achievement_id": id,
			"document_key":   documentKey,
			"error":          err.Error(),
		}).Error("Failed to update achievement in Couchbase")
		return fmt.Errorf("failed to update achievement: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_update",
		"achievement_id": id,
		"document_key":   documentKey,
	}).Info("Successfully updated achievement in Couchbase")

	return nil
}

func (r *AchievementRepository) DeleteAchievement(id string) error {
	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_delete",
		"achievement_id": id,
	}).Info("Deleting achievement from Couchbase")

	documentKey := fmt.Sprintf("achievement::%s", id)
	
	_, err := r.collection.Remove(documentKey, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":      "achievement_delete",
			"achievement_id": id,
			"document_key":   documentKey,
			"error":          err.Error(),
		}).Error("Failed to delete achievement from Couchbase")
		return fmt.Errorf("failed to delete achievement: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":      "achievement_delete",
		"achievement_id": id,
		"document_key":   documentKey,
	}).Info("Successfully deleted achievement from Couchbase")

	return nil
}

func (r *AchievementRepository) ListAchievements(limit, offset int) ([]*Achievement, error) {
	r.logger.WithFields(logrus.Fields{
		"operation": "achievement_list",
		"limit":     limit,
		"offset":    offset,
	}).Info("Listing achievements from Couchbase")

	if limit == 0 {
		limit = 50 // Default limit
	}

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "achievement" LIMIT %d OFFSET %d`,
		limit, offset,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for achievements")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "achievement_list",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for achievements")
		return nil, fmt.Errorf("failed to query achievements: %w", err)
	}
	defer result.Close()

	var achievements []*Achievement
	for result.Next() {
		var achievement Achievement
		if err := result.Row(&achievement); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation": "achievement_list",
				"error":     err.Error(),
			}).Error("Failed to decode achievement row")
			continue
		}
		achievements = append(achievements, &achievement)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "achievement_list",
			"error":     err.Error(),
		}).Error("Error during achievement query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation": "achievement_list",
		"count":     len(achievements),
		"limit":     limit,
		"offset":    offset,
	}).Info("Successfully retrieved achievements from Couchbase")

	return achievements, nil
}

func (r *AchievementRepository) GetAchievementsByStudent(studentID string) ([]*Achievement, error) {
	r.logger.WithFields(logrus.Fields{
		"operation":  "achievement_get_by_student",
		"student_id": studentID,
	}).Info("Retrieving achievements by student from Couchbase")

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "achievement" AND a.student_id = "%s"`,
		studentID,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for achievements by student")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":  "achievement_get_by_student",
			"student_id": studentID,
			"query":      query,
			"error":      err.Error(),
		}).Error("Failed to execute N1QL query for achievements by student")
		return nil, fmt.Errorf("failed to query achievements by student: %w", err)
	}
	defer result.Close()

	var achievements []*Achievement
	for result.Next() {
		var achievement Achievement
		if err := result.Row(&achievement); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation":  "achievement_get_by_student",
				"student_id": studentID,
				"error":      err.Error(),
			}).Error("Failed to decode achievement row")
			continue
		}
		achievements = append(achievements, &achievement)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":  "achievement_get_by_student",
			"student_id": studentID,
			"error":      err.Error(),
		}).Error("Error during achievement query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":  "achievement_get_by_student",
		"student_id": studentID,
		"count":      len(achievements),
	}).Info("Successfully retrieved achievements by student from Couchbase")

	return achievements, nil
}

func (r *AchievementRepository) GetAchievementsByCategory(category string) ([]*Achievement, error) {
	r.logger.WithFields(logrus.Fields{
		"operation": "achievement_get_by_category",
		"category":  category,
	}).Info("Retrieving achievements by category from Couchbase")

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "achievement" AND a.category = "%s"`,
		category,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for achievements by category")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "achievement_get_by_category",
			"category":  category,
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for achievements by category")
		return nil, fmt.Errorf("failed to query achievements by category: %w", err)
	}
	defer result.Close()

	var achievements []*Achievement
	for result.Next() {
		var achievement Achievement
		if err := result.Row(&achievement); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation": "achievement_get_by_category",
				"category":  category,
				"error":     err.Error(),
			}).Error("Failed to decode achievement row")
			continue
		}
		achievements = append(achievements, &achievement)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "achievement_get_by_category",
			"category":  category,
			"error":     err.Error(),
		}).Error("Error during achievement query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation": "achievement_get_by_category",
		"category":  category,
		"count":     len(achievements),
	}).Info("Successfully retrieved achievements by category from Couchbase")

	return achievements, nil
}

// Award CRUD operations
func (r *AchievementRepository) CreateAward(award *Award) error {
	r.logger.WithFields(logrus.Fields{
		"operation": "award_create",
		"award_id":  award.ID,
		"name":      award.Name,
	}).Info("Creating new award in Couchbase")

	// Set metadata
	award.CreatedAt = time.Now()
	award.UpdatedAt = time.Now()
	award.Type = "award"

	// Generate ID if not provided
	if award.ID == "" {
		award.ID = fmt.Sprintf("award_%d", time.Now().UnixNano())
		r.logger.WithField("generated_id", award.ID).Info("Generated new award ID")
	}

	documentKey := fmt.Sprintf("award::%s", award.ID)
	
	_, err := r.collection.Insert(documentKey, award, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "award_create",
			"award_id":     award.ID,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to create award in Couchbase")
		return fmt.Errorf("failed to create award: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "award_create",
		"award_id":     award.ID,
		"document_key": documentKey,
	}).Info("Successfully created award in Couchbase")

	return nil
}

func (r *AchievementRepository) GetAwardByID(id string) (*Award, error) {
	r.logger.WithFields(logrus.Fields{
		"operation": "award_get_by_id",
		"award_id":  id,
	}).Info("Retrieving award from Couchbase")

	documentKey := fmt.Sprintf("award::%s", id)
	
	result, err := r.collection.Get(documentKey, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "award_get_by_id",
			"award_id":     id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to retrieve award from Couchbase")
		return nil, fmt.Errorf("failed to get award: %w", err)
	}

	var award Award
	if err := result.Content(&award); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "award_get_by_id",
			"award_id":     id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to decode award document")
		return nil, fmt.Errorf("failed to decode award: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "award_get_by_id",
		"award_id":     id,
		"document_key": documentKey,
	}).Info("Successfully retrieved award from Couchbase")

	return &award, nil
}

func (r *AchievementRepository) ListAwards(limit, offset int) ([]*Award, error) {
	r.logger.WithFields(logrus.Fields{
		"operation": "award_list",
		"limit":     limit,
		"offset":    offset,
	}).Info("Listing awards from Couchbase")

	if limit == 0 {
		limit = 50 // Default limit
	}

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "award" LIMIT %d OFFSET %d`,
		limit, offset,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for awards")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "award_list",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for awards")
		return nil, fmt.Errorf("failed to query awards: %w", err)
	}
	defer result.Close()

	var awards []*Award
	for result.Next() {
		var award Award
		if err := result.Row(&award); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation": "award_list",
				"error":     err.Error(),
			}).Error("Failed to decode award row")
			continue
		}
		awards = append(awards, &award)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "award_list",
			"error":     err.Error(),
		}).Error("Error during award query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation": "award_list",
		"count":     len(awards),
		"limit":     limit,
		"offset":    offset,
	}).Info("Successfully retrieved awards from Couchbase")

	return awards, nil
}
