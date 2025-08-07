package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"schoolmgmt/services/academic-service/internal/models"
	"schoolmgmt/services/academic-service/internal/service"
	"schoolmgmt/shared/pkg/response"
)

type AcademicHandler struct {
	service *service.AcademicService
}

func NewAcademicHandler(service *service.AcademicService) *AcademicHandler {
	log.Println("AcademicHandler: Creating new academic handler with Couchbase-backed service")
	return &AcademicHandler{
		service: service,
	}
}

func (h *AcademicHandler) CreateAcademic(c *gin.Context) {
	log.Println("AcademicHandler: CreateAcademic called")
	var req models.CreateAcademicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("AcademicHandler: Error binding JSON: %v", err)
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	academic := &models.Academic{
		ID:            "academic-" + uuid.New().String(),
		StudentID:     req.StudentID,
		TeacherID:     req.TeacherID,
		Subject:       req.Subject,
		Grade:         req.Grade,
		Semester:      req.Semester,
		AcademicYear:  req.AcademicYear,
		ExamType:      req.ExamType,
		MaxMarks:      req.MaxMarks,
		ObtainedMarks: req.ObtainedMarks,
		Percentage:    (req.ObtainedMarks / req.MaxMarks) * 100,
		Status:        "pass",
		Type:          "academic",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if academic.Percentage < 40 {
		academic.Status = "fail"
	}

	log.Printf("AcademicHandler: Creating academic record for student %s", academic.StudentID)
	ctx := context.Background()
	if err := h.service.CreateAcademic(ctx, academic); err != nil {
		log.Printf("AcademicHandler: Error creating academic record: %v", err)
		response.InternalServerError(c, "Failed to create academic record: "+err.Error())
		return
	}

	log.Printf("AcademicHandler: Successfully created academic record %s", academic.ID)
	response.Created(c, academic, "Academic record created successfully")
}

func (h *AcademicHandler) GetAcademic(c *gin.Context) {
	id := c.Param("id")
	log.Printf("AcademicHandler: GetAcademic called for ID: %s", id)

	ctx := context.Background()
	academic, err := h.service.GetAcademic(ctx, id)
	if err != nil {
		log.Printf("AcademicHandler: Error retrieving academic record %s: %v", id, err)
		response.NotFound(c, "Academic record not found")
		return
	}

	log.Printf("AcademicHandler: Successfully retrieved academic record %s", id)
	response.Success(c, academic, "Academic record retrieved successfully")
}

func (h *AcademicHandler) UpdateAcademic(c *gin.Context) {
	id := c.Param("id")
	log.Printf("AcademicHandler: UpdateAcademic called for ID: %s", id)

	var req models.UpdateAcademicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("AcademicHandler: Error binding JSON: %v", err)
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	ctx := context.Background()
	academic, err := h.service.GetAcademic(ctx, id)
	if err != nil {
		log.Printf("AcademicHandler: Academic record %s not found for update: %v", id, err)
		response.NotFound(c, "Academic record not found")
		return
	}

	// Update fields
	if req.Subject != "" {
		academic.Subject = req.Subject
	}
	if req.Grade != "" {
		academic.Grade = req.Grade
	}
	if req.Semester != "" {
		academic.Semester = req.Semester
	}
	if req.AcademicYear != "" {
		academic.AcademicYear = req.AcademicYear
	}
	if req.ExamType != "" {
		academic.ExamType = req.ExamType
	}
	if req.MaxMarks > 0 {
		academic.MaxMarks = req.MaxMarks
	}
	if req.ObtainedMarks >= 0 {
		academic.ObtainedMarks = req.ObtainedMarks
	}

	// Recalculate percentage and status
	if academic.MaxMarks > 0 {
		academic.Percentage = (academic.ObtainedMarks / academic.MaxMarks) * 100
		academic.Status = "pass"
		if academic.Percentage < 40 {
			academic.Status = "fail"
		}
	}

	academic.UpdatedAt = time.Now()

	if err := h.service.UpdateAcademic(ctx, academic); err != nil {
		log.Printf("AcademicHandler: Error updating academic record %s: %v", id, err)
		response.InternalServerError(c, "Failed to update academic record: "+err.Error())
		return
	}

	log.Printf("AcademicHandler: Successfully updated academic record %s", id)
	response.Success(c, academic, "Academic record updated successfully")
}

func (h *AcademicHandler) DeleteAcademic(c *gin.Context) {
	id := c.Param("id")
	log.Printf("AcademicHandler: DeleteAcademic called for ID: %s", id)

	ctx := context.Background()
	if err := h.service.DeleteAcademic(ctx, id); err != nil {
		log.Printf("AcademicHandler: Error deleting academic record %s: %v", id, err)
		response.NotFound(c, "Academic record not found")
		return
	}

	log.Printf("AcademicHandler: Successfully deleted academic record %s", id)
	response.Success(c, nil, "Academic record deleted successfully")
}

func (h *AcademicHandler) ListAcademics(c *gin.Context) {
	log.Println("AcademicHandler: ListAcademics called")

	ctx := context.Background()
	academics, err := h.service.GetAllAcademics(ctx)
	if err != nil {
		log.Printf("AcademicHandler: Error retrieving academics: %v", err)
		response.InternalServerError(c, "Failed to retrieve academics: "+err.Error())
		return
	}

	log.Printf("AcademicHandler: Successfully retrieved %d academic records", len(academics))
	response.Success(c, academics, "Academic records retrieved successfully")
}

func (h *AcademicHandler) CreateClass(c *gin.Context) {
	log.Println("AcademicHandler: CreateClass called")
	var req models.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("AcademicHandler: Error binding JSON: %v", err)
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	class := &models.Class{
		ID:        "class-" + uuid.New().String(),
		Name:      req.Name,
		Subject:   req.Subject,
		TeacherID: req.TeacherID,
		Students:  req.Students,
		Schedule:  req.Schedule,
		Type:      "class",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Printf("AcademicHandler: Creating class %s", class.Name)
	ctx := context.Background()
	if err := h.service.CreateClass(ctx, class); err != nil {
		log.Printf("AcademicHandler: Error creating class: %v", err)
		response.InternalServerError(c, "Failed to create class: "+err.Error())
		return
	}

	log.Printf("AcademicHandler: Successfully created class %s", class.ID)
	response.Created(c, class, "Class created successfully")
}

func (h *AcademicHandler) GetClass(c *gin.Context) {
	id := c.Param("id")
	log.Printf("AcademicHandler: GetClass called for ID: %s", id)

	ctx := context.Background()
	class, err := h.service.GetClass(ctx, id)
	if err != nil {
		log.Printf("AcademicHandler: Error retrieving class %s: %v", id, err)
		response.NotFound(c, "Class not found")
		return
	}

	log.Printf("AcademicHandler: Successfully retrieved class %s", id)
	response.Success(c, class, "Class retrieved successfully")
}

func (h *AcademicHandler) ListClasses(c *gin.Context) {
	log.Println("AcademicHandler: ListClasses called")

	ctx := context.Background()
	classes, err := h.service.GetAllClasses(ctx)
	if err != nil {
		log.Printf("AcademicHandler: Error retrieving classes: %v", err)
		response.InternalServerError(c, "Failed to retrieve classes: "+err.Error())
		return
	}

	log.Printf("AcademicHandler: Successfully retrieved %d classes", len(classes))
	response.Success(c, classes, "Classes retrieved successfully")
}

func (h *AcademicHandler) HealthCheck(c *gin.Context) {
	log.Println("AcademicHandler: HealthCheck called")
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "academic-service",
		"time":    time.Now().Format(time.RFC3339),
	})
}
