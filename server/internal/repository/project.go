package repository

import "bocchikitsunei_webportfolio/internal/entities"

type ProjectRepository interface {
	GetAllProject() ([]entities.Project, error)
	GetProjectById(int) (*entities.Project, error)

	//////////////////////////////////////////////////////////////

	PostAddProject(project *entities.Project) error
	GetProjectsFirstFour() ([]entities.Project, error)

	GetEditProject(int) (*entities.Project, error)
	UpdateEditProject(project *entities.Project) error

	DeleteProject(projectID int) error
}
