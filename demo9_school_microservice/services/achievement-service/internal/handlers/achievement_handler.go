package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"schoolmgmt/services/achievement-service/internal/models"
	"schoolmgmt/shared/pkg/response"
)

type AchievementHandler struct {
	achievements map[string]*models.Achievement
	badges       map[string]*models.Badge
}

func NewAchievementHandler(service interface{}) *AchievementHandler {
	return &AchievementHandler{
		achievements: make(map[string]*models.Achievement),
		badges:       make(map[string]*models.Badge),
	}
}

func (h *AchievementHandler) CreateAchievement(c *gin.Context) {
	var req models.CreateAchievementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	achievement := &models.Achievement{
		ID:           "achievement-" + strconv.Itoa(len(h.achievements)+1),
		StudentID:    req.StudentID,
		TeacherID:    req.TeacherID,
		Title:        req.Title,
		Description:  req.Description,
		Category:     req.Category,
		Level:        req.Level,
		AwardedDate:  req.AwardedDate,
		AcademicYear: req.AcademicYear,
		Semester:     req.Semester,
		Points:       req.Points,
		Certificate:  req.Certificate,
		Verified:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Type:         "achievement",
	}

	h.achievements[achievement.ID] = achievement
	response.Created(c, achievement, "Achievement created successfully")
}

func (h *AchievementHandler) GetAchievement(c *gin.Context) {
	id := c.Param("id")
	achievement, exists := h.achievements[id]
	if !exists {
		response.NotFound(c, "Achievement not found")
		return
	}

	response.Success(c, achievement, "Achievement retrieved successfully")
}

func (h *AchievementHandler) UpdateAchievement(c *gin.Context) {
	id := c.Param("id")
	achievement, exists := h.achievements[id]
	if !exists {
		response.NotFound(c, "Achievement not found")
		return
	}

	var req models.UpdateAchievementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	if req.Title != nil {
		achievement.Title = *req.Title
	}
	if req.Verified != nil {
		achievement.Verified = *req.Verified
	}
	if req.VerifiedBy != nil {
		achievement.VerifiedBy = *req.VerifiedBy
	}

	achievement.UpdatedAt = time.Now()
	response.Success(c, achievement, "Achievement updated successfully")
}

func (h *AchievementHandler) DeleteAchievement(c *gin.Context) {
	id := c.Param("id")
	if _, exists := h.achievements[id]; !exists {
		response.NotFound(c, "Achievement not found")
		return
	}

	delete(h.achievements, id)
	response.Success(c, nil, "Achievement deleted successfully")
}

func (h *AchievementHandler) ListAchievements(c *gin.Context) {
	var achievements []*models.Achievement
	for _, achievement := range h.achievements {
		achievements = append(achievements, achievement)
	}

	response.Success(c, map[string]interface{}{
		"achievements": achievements,
		"count":        len(achievements),
	}, "Achievements retrieved successfully")
}

func (h *AchievementHandler) CreateBadge(c *gin.Context) {
	var req models.CreateBadgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	badge := &models.Badge{
		ID:         "badge-" + strconv.Itoa(len(h.badges)+1),
		StudentID:  req.StudentID,
		BadgeName:  req.BadgeName,
		BadgeType:  req.BadgeType,
		Criteria:   req.Criteria,
		EarnedDate: time.Now(),
		ValidUntil: req.ValidUntil,
		BadgeIcon:  req.BadgeIcon,
		Points:     req.Points,
		IsActive:   true,
		CreatedAt:  time.Now(),
		Type:       "badge",
	}

	h.badges[badge.ID] = badge
	response.Created(c, badge, "Badge created successfully")
}

func (h *AchievementHandler) GetBadge(c *gin.Context) {
	id := c.Param("id")
	badge, exists := h.badges[id]
	if !exists {
		response.NotFound(c, "Badge not found")
		return
	}

	response.Success(c, badge, "Badge retrieved successfully")
}

func (h *AchievementHandler) ListBadges(c *gin.Context) {
	var badges []*models.Badge
	for _, badge := range h.badges {
		badges = append(badges, badge)
	}

	response.Success(c, map[string]interface{}{
		"badges": badges,
		"count":  len(badges),
	}, "Badges retrieved successfully")
}

func (h *AchievementHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "achievement-service",
	})
}
