package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	var req models.CreateAcademicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	academic := &models.Academic{
		ID:            "academic-" + strconv.Itoa(len(h.academics)+1),
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
	}

	if academic.Percentage < 40 {
		academic.Status = "fail"
	}

	h.academics[academic.ID] = academic
	response.Created(c, academic, "Academic record created successfully")
}

func (h *AcademicHandler) GetAcademic(c *gin.Context) {
	id := c.Param("id")
	academic, exists := h.academics[id]
	if !exists {
		response.NotFound(c, "Academic record not found")
		return
	}

	response.Success(c, academic, "Academic record retrieved successfully")
}

func (h *AcademicHandler) UpdateAcademic(c *gin.Context) {
	id := c.Param("id")
	academic, exists := h.academics[id]
	if !exists {
		response.NotFound(c, "Academic record not found")
		return
	}

	var req models.UpdateAcademicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	if req.ObtainedMarks != nil {
		academic.ObtainedMarks = *req.ObtainedMarks
		academic.Percentage = (academic.ObtainedMarks / academic.MaxMarks) * 100
		if academic.Percentage < 40 {
			academic.Status = "fail"
		} else {
			academic.Status = "pass"
		}
	}

	response.Success(c, academic, "Academic record updated successfully")
}

func (h *AcademicHandler) DeleteAcademic(c *gin.Context) {
	id := c.Param("id")
	if _, exists := h.academics[id]; !exists {
		response.NotFound(c, "Academic record not found")
		return
	}

	delete(h.academics, id)
	response.Success(c, nil, "Academic record deleted successfully")
}

func (h *AcademicHandler) ListAcademics(c *gin.Context) {
	var academics []*models.Academic
	for _, academic := range h.academics {
		academics = append(academics, academic)
	}

	response.Success(c, map[string]interface{}{
		"academics": academics,
		"count":     len(academics),
	}, "Academic records retrieved successfully")
}

func (h *AcademicHandler) CreateClass(c *gin.Context) {
	var req models.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	class := &models.Class{
		ID:           "class-" + strconv.Itoa(len(h.classes)+1),
		ClassName:    req.ClassName,
		Grade:        req.Grade,
		Section:      req.Section,
		TeacherID:    req.TeacherID,
		Subject:      req.Subject,
		AcademicYear: req.AcademicYear,
		Semester:     req.Semester,
		StudentIDs:   req.StudentIDs,
		Schedule:     req.Schedule,
		MaxCapacity:  req.MaxCapacity,
		Status:       "active",
		Type:         "class",
	}

	h.classes[class.ID] = class
	response.Created(c, class, "Class created successfully")
}

func (h *AcademicHandler) GetClass(c *gin.Context) {
	id := c.Param("id")
	class, exists := h.classes[id]
	if !exists {
		response.NotFound(c, "Class not found")
		return
	}

	response.Success(c, class, "Class retrieved successfully")
}

func (h *AcademicHandler) ListClasses(c *gin.Context) {
	var classes []*models.Class
	for _, class := range h.classes {
		classes = append(classes, class)
	}

	response.Success(c, map[string]interface{}{
		"classes": classes,
		"count":   len(classes),
	}, "Classes retrieved successfully")
}

func (h *AcademicHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "academic-service",
	})
}
