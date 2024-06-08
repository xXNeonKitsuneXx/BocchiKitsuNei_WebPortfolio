package handler

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

type projectHandler struct {
	projectSer service.ProjectService
	jwtSecret  string
	uploadSer  service.UploadService
}

func NewProjectHandler(projectSer service.ProjectService, jwtSecret string, uploadSer service.UploadService) projectHandler {
	return projectHandler{projectSer: projectSer, jwtSecret: jwtSecret, uploadSer: uploadSer}
}

func (h *projectHandler) GetProjects(c *fiber.Ctx) error {
	projectsResponse := make([]dtos.ProjectsResponse, 0)

	projects, err := h.projectSer.GetProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		projectsResponse = append(projectsResponse, dtos.ProjectsResponse{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		})
	}
	return c.JSON(projectsResponse)
}

func (h *projectHandler) GetProjectById(c *fiber.Ctx) error {

	projectIDReceive, err := strconv.Atoi(c.Params("ProjectID"))

	project, err := h.projectSer.GetProjectId(projectIDReceive)
	if err != nil {
		return err
	}

	projectResponse := dtos.ProjectsResponse{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) AddProject(c *fiber.Ctx) error {
	//// Extract the token from the request headers
	//token := c.Get("Authorization")
	//
	//// Check if the token is empty
	//if token == "" {
	//	return errors.New("token is missing")
	//}
	//
	//// Extract the user ID from the token
	//userIDExtract, err := utils.ExtractUserIDFromToken(strings.Replace(token, "Bearer ", "", 1), h.jwtSecret)
	//if err != nil {
	//	return err
	//}

	userIDExtract, err := strconv.Atoi(c.Params("UserID"))

	var request dtos.AddProjectRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	// Check if a file is uploaded
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "File not found")
	}

	// Call upload service to upload the file
	fileURL, err := h.uploadSer.UploadFile(file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload file")
	}

	// Set the uploaded file URL in the registration request
	request.ProjectPicture = fileURL

	// Check if project_picture field is empty or nil
	if request.ProjectPicture == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Project picture is required")
	}

	project, err := h.projectSer.CreateAddProject(userIDExtract, request)
	if err != nil {
		return err
	}

	projectResponse := dtos.AddProjectRequest{
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) GetProjectsFirstFour(c *fiber.Ctx) error {
	projectsResponse := make([]dtos.ProjectsFirstFourResponse, 0)

	projects, err := h.projectSer.GetProjectsFirstFour()
	if err != nil {
		return err
	}

	for _, project := range projects {
		projectsResponse = append(projectsResponse, dtos.ProjectsFirstFourResponse{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		})
	}
	return c.JSON(projectsResponse)
}

func (h *projectHandler) GetEditProject(c *fiber.Ctx) error {
	projectIDReceive, err := strconv.Atoi(c.Params("ProjectID"))

	project, err := h.projectSer.GetEditProject(projectIDReceive)
	if err != nil {
		return err
	}

	projectResponse := dtos.ProjectsResponse{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) UpdateEditProject(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("ProjectID"))
	if err != nil {
		return err
	}

	var req dtos.EditProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	project, err := h.projectSer.UpdateEditProject(projectID, req)
	if err != nil {
		return err
	}

	projectResponse := dtos.EditProjectRequest{
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) DeleteProject(c *fiber.Ctx) error {
	projectIDReceive, err := strconv.Atoi(c.Params("ProjectID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid projectID"})
	}

	err = h.projectSer.DeleteProject(projectIDReceive)
	if err != nil {
		if strings.Contains(err.Error(), "project not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "project not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()}) //failed to delete project
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "project deleted successfully"})
}
