package repositories

import "hollyways/models"

type ProjectRepository interface {
	CreateProject(project models.Project) error
	FindProject() ([]models.Project, error)
	GetProject(id int) (models.Project, error)
	UpdateProjectByAdmin(project models.Project) error
}

func (r *repository) CreateProject(project models.Project) error {
	err := r.db.Create(&project).Error

	return err
}

func (r *repository) FindProject() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Find(&projects).Error

	return projects, err
}

func (r *repository) GetProject(id int) (models.Project, error) {
	var project models.Project
	err := r.db.First(&project, "id = ?", id).Error

	return project, err
}

func (r *repository) UpdateProjectByAdmin(project models.Project) error {
	err := r.db.Save(&project).Error

	return err
}
