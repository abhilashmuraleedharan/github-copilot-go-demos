package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"schoolmgmt/services/teacher-service/internal/models"
	"schoolmgmt/shared/pkg/response"
)

type TeacherHandler struct {
	// For demo purposes, we'll use a simple in-memory store
	// In real implementation, this would use a service layer
	teachers map[string]*models.Teacher
}

func NewTeacherHandler(service interface{}) *TeacherHandler {
	return &TeacherHandler{
		teachers: make(map[string]*models.Teacher),
	}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var req models.CreateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// For demo - simplified creation
	teacher := &models.Teacher{
		ID:        "teacher-" + strconv.Itoa(len(h.teachers)+1),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Department: req.Department,
		Status:    "active",
		Type:      "teacher",
	}

	h.teachers[teacher.ID] = teacher
	response.Created(c, teacher, "Teacher created successfully")
}

func (h *TeacherHandler) GetTeacher(c *gin.Context) {
	id := c.Param("id")
	teacher, exists := h.teachers[id]
	if !exists {
		response.NotFound(c, "Teacher not found")
		return
	}

	response.Success(c, teacher, "Teacher retrieved successfully")
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	teacher, exists := h.teachers[id]
	if !exists {
		response.NotFound(c, "Teacher not found")
		return
	}

	var req models.UpdateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	// Update fields if provided
	if req.FirstName != nil {
		teacher.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		teacher.LastName = *req.LastName
	}
	if req.Email != nil {
		teacher.Email = *req.Email
	}
	if req.Status != nil {
		teacher.Status = *req.Status
	}

	response.Success(c, teacher, "Teacher updated successfully")
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if _, exists := h.teachers[id]; !exists {
		response.NotFound(c, "Teacher not found")
		return
	}

	delete(h.teachers, id)
	response.Success(c, nil, "Teacher deleted successfully")
}

func (h *TeacherHandler) ListTeachers(c *gin.Context) {
	var teachers []*models.Teacher
	for _, teacher := range h.teachers {
		teachers = append(teachers, teacher)
	}

	response.Success(c, map[string]interface{}{
		"teachers": teachers,
		"count":    len(teachers),
	}, "Teachers retrieved successfully")
}

func (h *TeacherHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "teacher-service",
	})
}
