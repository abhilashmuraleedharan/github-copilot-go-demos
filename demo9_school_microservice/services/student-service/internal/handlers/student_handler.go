package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"schoolmgmt/services/student-service/internal/models"
	"schoolmgmt/services/student-service/internal/service"
	"schoolmgmt/shared/pkg/response"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// @Summary Create a new student
// @Description Create a new student record
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.CreateStudentRequest true "Student data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /students [post]
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req models.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	student, err := h.service.CreateStudent(&req)
	if err != nil {
		if err.Error() == "student with email "+req.Email+" already exists" {
			response.Conflict(c, err.Error())
			return
		}
		response.InternalError(c, "Failed to create student: "+err.Error())
		return
	}

	response.Created(c, student, "Student created successfully")
}

// @Summary Get a student by ID
// @Description Get a student record by ID
// @Tags students
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /students/{id} [get]
func (h *StudentHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Student ID is required")
		return
	}

	student, err := h.service.GetStudent(id)
	if err != nil {
		if err.Error() == "student not found" {
			response.NotFound(c, "Student not found")
			return
		}
		response.InternalError(c, "Failed to get student: "+err.Error())
		return
	}

	response.Success(c, student, "Student retrieved successfully")
}

// @Summary Update a student
// @Description Update a student record
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Param student body models.UpdateStudentRequest true "Student update data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /students/{id} [put]
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Student ID is required")
		return
	}

	var req models.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	student, err := h.service.UpdateStudent(id, &req)
	if err != nil {
		if err.Error() == "student not found" {
			response.NotFound(c, "Student not found")
			return
		}
		response.InternalError(c, "Failed to update student: "+err.Error())
		return
	}

	response.Success(c, student, "Student updated successfully")
}

// @Summary Delete a student
// @Description Delete a student record
// @Tags students
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /students/{id} [delete]
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Student ID is required")
		return
	}

	err := h.service.DeleteStudent(id)
	if err != nil {
		if err.Error() == "student not found" {
			response.NotFound(c, "Student not found")
			return
		}
		response.InternalError(c, "Failed to delete student: "+err.Error())
		return
	}

	response.Success(c, nil, "Student deleted successfully")
}

// @Summary List students
// @Description Get a list of students with pagination
// @Tags students
// @Produce json
// @Param limit query int false "Number of students to return (default 10, max 100)"
// @Param offset query int false "Number of students to skip (default 0)"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /students [get]
func (h *StudentHandler) ListStudents(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.BadRequest(c, "Invalid limit parameter")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.BadRequest(c, "Invalid offset parameter")
		return
	}

	students, err := h.service.ListStudents(limit, offset)
	if err != nil {
		response.InternalError(c, "Failed to list students: "+err.Error())
		return
	}

	response.Success(c, map[string]interface{}{
		"students": students,
		"count":    len(students),
		"limit":    limit,
		"offset":   offset,
	}, "Students retrieved successfully")
}

// @Summary Health check
// @Description Health check endpoint
// @Tags health
// @Produce json
// @Success 200 {object} response.APIResponse
// @Router /health [get]
func (h *StudentHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "student-service",
	})
}
