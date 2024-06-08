package repository

import (
	"bocchikitsunei_webportfolio/internal/entities"

	"gorm.io/gorm"
)

type projectRepositoryDB struct {
	db *gorm.DB
}

func NewProjectRepositoryDB(db *gorm.DB) projectRepositoryDB {
	return projectRepositoryDB{db: db}
}

func (r projectRepositoryDB) GetAllProject() ([]entities.Project, error) {
	projects := []entities.Project{}
	result := r.db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r projectRepositoryDB) GetProjectById(projectid int) (*entities.Project, error) {
	projects := entities.Project{}
	result := r.db.Where("project_id = ?", projectid).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}

func (r projectRepositoryDB) PostAddProject(project *entities.Project) error {
	result := r.db.Create(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r projectRepositoryDB) GetProjectsFirstFour() ([]entities.Project, error) {
	projects := []entities.Project{}
	result := r.db.Limit(4).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r projectRepositoryDB) GetEditProject(projectid int) (*entities.Project, error) {
	projects := entities.Project{}
	result := r.db.Where("project_id = ?", projectid).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}

func (r projectRepositoryDB) UpdateEditProject(project *entities.Project) error {
	result := r.db.Updates(project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r projectRepositoryDB) DeleteProject(projectid int) error {
	projects := entities.Project{}
	result := r.db.Where("project_id = ?", projectid).Unscoped().Delete(&projects)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
