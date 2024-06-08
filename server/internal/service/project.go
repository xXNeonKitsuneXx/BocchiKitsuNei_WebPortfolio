package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
)

type ProjectService interface {
	GetProjects() ([]entities.Project, error)
	GetProjectId(int) (*entities.Project, error)
	///////////////////////////////////////////////////////////

	CreateAddProject(int, dtos.AddProjectRequest) (*entities.Project, error)
	GetProjectsFirstFour() ([]entities.Project, error)

	GetEditProject(int) (*entities.Project, error)
	UpdateEditProject(int, dtos.EditProjectRequest) (*entities.Project, error)

	DeleteProject(projectID int) error
}
