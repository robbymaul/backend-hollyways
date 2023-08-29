package handlers

import (
	"fmt"
	projectdto "hollyways/dto/project"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	"hollyways/repositories"
	"hollyways/utilities"
	"html"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type projectHandler struct {
	ProjectRepository repositories.ProjectRepository
}

func ProjectHandler(ProjectRepository repositories.ProjectRepository) *projectHandler {
	return &projectHandler{ProjectRepository}
}

func (h *projectHandler) CreateProject(c *gin.Context) {
	var err error

	userLogin, _ := c.Get("userLogin")
	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))
	if userRole != 1 && userRole != 2 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	dataFile, _ := c.Get("file")
	fmt.Println("ini data file", dataFile)
	projectName := c.PostForm("projectName")
	projectDescription := c.PostForm("projectDescription")
	targetDonation, _ := strconv.Atoi(c.PostForm("target"))
	startDate, _ := utilities.ParseTime(c.PostForm("startDate"))
	dueDate, _ := utilities.ParseTime(c.PostForm("dueDate"))

	request := projectdto.ProjectRequestDTO{
		ProjectName:        projectName,
		ProjectDescription: projectDescription,
		ProjectImage:       dataFile.(string),
		TargetDonation:     targetDonation,
		StartDate:          startDate,
		DueDate:            dueDate,
	}

	validation := validator.New()
	if err = validation.Struct(request); err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Validation error",
		})
		return
	}

	project := models.Project{
		ProjectName:        html.EscapeString(request.ProjectName),
		ProjectDescription: html.EscapeString(request.ProjectDescription),
		ProjectImage:       request.ProjectImage,
		Donation:           0,
		TargetDonation:     request.TargetDonation,
		StartDate:          request.StartDate,
		DueDate:            request.DueDate,
		Progress:           0,
	}

	err = h.ProjectRepository.CreateProject(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create project donation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Create project successfully",
	})
}

func (h *projectHandler) FindProject(c *gin.Context) {
	var err error
	var projectsDTO []projectdto.ProjectResponseDTO

	projects, err := h.ProjectRepository.FindProject()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(projects) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "No project data found",
			"data":    projects,
		})
		return
	}

	for _, project := range projects {
		projectsDTO = append(projectsDTO, ConvertProjectResponse(project))
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   projectsDTO,
	})

}

func (h *projectHandler) GetProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	project, err := h.ProjectRepository.GetProject(projectId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, dtoResult.ErrorResult{
				Status:  http.StatusNotFound,
				Message: "Data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if project.ProjectName == "" {
		c.JSON(http.StatusNotFound, dtoResult.ErrorResult{
			Status:  http.StatusNotFound,
			Message: "Project data not found",
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   ConvertProjectResponse(project),
	})
}

func (h *projectHandler) UpdateProjectByAdmin(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")
	dataFile, _ := c.Get("file")

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))
	if userRole != 1 && userRole != 2 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	dueDate, _ := utilities.ParseTime(c.PostForm("dueDate"))

	request := projectdto.ProjectUpdateRequestDTO{
		ProjectName:        c.PostForm("projectName"),
		ProjectDescription: c.PostForm("projectDescription"),
		ProjectImage:       dataFile.(string),
		DueDate:            dueDate,
	}

	project, err := h.ProjectRepository.GetProject(projectId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, dtoResult.ErrorResult{
				Status:  http.StatusNotFound,
				Message: "Data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if request.ProjectName != "" {
		project.ProjectName = request.ProjectName
	}

	if request.ProjectDescription != "" {
		project.ProjectDescription = request.ProjectDescription
	}

	if request.ProjectImage != "" {
		project.ProjectImage = request.ProjectImage
	}

	if !request.DueDate.IsZero() {
		project.DueDate = request.DueDate
	}

	err = h.ProjectRepository.UpdateProjectByAdmin(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Update project successfully",
	})

}

// function for handle logic delet project by admin (with gin.context)
func (h *projectHandler) DeleteProjectByAdmin(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")
	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))
	if userRole != 1 && userRole != 2 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	project, err := h.ProjectRepository.GetProject(projectId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "id not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	err = h.ProjectRepository.DeleteProjectByAdmin(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to delete project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Delete project successfully",
	})
}

// function convert response model project with data transfer objct project response
func ConvertProjectResponse(p models.Project) projectdto.ProjectResponseDTO {
	return projectdto.ProjectResponseDTO{
		ID:                 int(p.ID),
		ProjectName:        p.ProjectName,
		ProjectDescription: p.ProjectDescription,
		ProjectImage:       p.ProjectImage,
		Donation:           p.Donation,
		TargetDonation:     p.TargetDonation,
		StartDate:          p.StartDate.Format("2006-01-02"),
		DueDate:            p.DueDate.Format("2006-01-02"),
		Progress:           float64(p.Progress),
	}
}
