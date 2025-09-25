// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package handler

import (
	"net/http"
	"school-microservice/internal/models"
	"school-microservice/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StudentHandler handles HTTP requests for students
type StudentHandler struct {
	service service.StudentService
}

// NewStudentHandler creates a new student handler
func NewStudentHandler(service service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// CreateStudent handles POST /api/v1/students
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req models.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	student, err := h.service.CreateStudent(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to create student",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Student created successfully",
		Data:    student,
	})
}

// GetStudent handles GET /api/v1/students/:id
func (h *StudentHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Student ID is required",
		})
		return
	}

	student, err := h.service.GetStudent(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Student not found",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Student retrieved successfully",
		Data:    student,
	})
}

// GetAllStudents handles GET /api/v1/students
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	students, total, err := h.service.GetAllStudents(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to retrieve students",
			Error:   err.Error(),
		})
		return
	}

	totalPages := (total + pageSize - 1) / pageSize

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Success:    true,
		Message:    "Students retrieved successfully",
		Data:       students,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}

// UpdateStudent handles PUT /api/v1/students/:id
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Student ID is required",
		})
		return
	}

	var req models.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	student, err := h.service.UpdateStudent(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to update student",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Student updated successfully",
		Data:    student,
	})
}

// DeleteStudent handles DELETE /api/v1/students/:id
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Student ID is required",
		})
		return
	}

	err := h.service.DeleteStudent(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to delete student",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Student deleted successfully",
	})
}

// GetStudentsByGrade handles GET /api/v1/students/grade/:grade
func (h *StudentHandler) GetStudentsByGrade(c *gin.Context) {
	grade := c.Param("grade")
	if grade == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Grade is required",
		})
		return
	}

	students, err := h.service.GetStudentsByGrade(c.Request.Context(), grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to retrieve students by grade",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Students retrieved successfully",
		Data:    students,
	})
}

// Handler aggregates all HTTP handlers
type Handler struct {
	Student     *StudentHandler
	Teacher     *TeacherHandler
	Class       *ClassHandler
	Academic    *AcademicHandler
	Achievement *AchievementHandler
}

// NewHandler creates a new handler instance
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Student:     NewStudentHandler(service.Student),
		Teacher:     NewTeacherHandler(service.Teacher),
		Class:       NewClassHandler(service.Class),
		Academic:    NewAcademicHandler(service.Academic),
		Achievement: NewAchievementHandler(service.Achievement),
	}
}

// Placeholder handler implementations for other entities
type TeacherHandler struct {
	service service.TeacherService
}

func NewTeacherHandler(service service.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Teacher endpoints not fully implemented",
	})
}

func (h *TeacherHandler) GetTeacher(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Teacher endpoints not fully implemented",
	})
}

type ClassHandler struct {
	service service.ClassService
}

func NewClassHandler(service service.ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Class endpoints not fully implemented",
	})
}

func (h *ClassHandler) GetClass(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Class endpoints not fully implemented",
	})
}

type AcademicHandler struct {
	service service.AcademicService
}

func NewAcademicHandler(service service.AcademicService) *AcademicHandler {
	return &AcademicHandler{service: service}
}

func (h *AcademicHandler) CreateAcademic(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Academic endpoints not fully implemented",
	})
}

func (h *AcademicHandler) GetAcademic(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Academic endpoints not fully implemented",
	})
}

type AchievementHandler struct {
	service service.AchievementService
}

func NewAchievementHandler(service service.AchievementService) *AchievementHandler {
	return &AchievementHandler{service: service}
}

func (h *AchievementHandler) CreateAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Achievement endpoints not fully implemented",
	})
}

func (h *AchievementHandler) GetAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Achievement endpoints not fully implemented",
	})
}