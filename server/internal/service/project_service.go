package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
	"bocchikitsunei_webportfolio/internal/repository"
	"bocchikitsunei_webportfolio/internal/utils/v"
	"log"
)

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) projectService {
	return projectService{projectRepo: projectRepo}
}

func (s projectService) GetProjects() ([]entities.Project, error) {
	projects, err := s.projectRepo.GetAllProject()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponses := []entities.Project{}
	for _, project := range projects {
		projectResponse := entities.Project{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		}
		projectResponses = append(projectResponses, projectResponse)
	}
	return projectResponses, nil
}

func (s projectService) GetProjectId(projectid int) (*entities.Project, error) {
	project, err := s.projectRepo.GetProjectById(projectid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponse := entities.Project{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}
	return &projectResponse, nil
}

func (s projectService) CreateAddProject(userid int, req dtos.AddProjectRequest) (*entities.Project, error) {
	project := &entities.Project{
		UserID:             v.UintPtr(userid),
		ProjectName:        req.ProjectName,
		ProjectTag:         req.ProjectTag,
		ProjectLanguage:    req.ProjectLanguage,
		ProjectGitHubURL:   req.ProjectGitHubURL,
		ProjectDescription: req.ProjectDescription,
		ProjectPicture:     req.ProjectPicture,
	}

	err := s.projectRepo.PostAddProject(project)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return project, nil
}

func (s projectService) GetProjectsFirstFour() ([]entities.Project, error) {
	projects, err := s.projectRepo.GetProjectsFirstFour()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponses := []entities.Project{}
	for _, project := range projects {
		projectResponse := entities.Project{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		}
		projectResponses = append(projectResponses, projectResponse)
	}
	return projectResponses, nil
}

func (s projectService) GetEditProject(projectid int) (*entities.Project, error) {
	project, err := s.projectRepo.GetEditProject(projectid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponse := entities.Project{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}
	return &projectResponse, nil
}

func (s projectService) UpdateEditProject(projectid int, req dtos.EditProjectRequest) (*entities.Project, error) {
	project := &entities.Project{
		ProjectID:          v.UintPtr(projectid),
		ProjectName:        req.ProjectName,
		ProjectTag:         req.ProjectTag,
		ProjectLanguage:    req.ProjectLanguage,
		ProjectGitHubURL:   req.ProjectGitHubURL,
		ProjectDescription: req.ProjectDescription,
		ProjectPicture:     req.ProjectPicture,
	}

	err := s.projectRepo.UpdateEditProject(project)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return project, nil
}
