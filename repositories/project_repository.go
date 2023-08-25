package repositories

import "hollyways/models"

// contract interface function query database model table projects
type ProjectRepository interface {
	CreateProject(project models.Project) error
	FindProject() ([]models.Project, error)
	GetProject(id int) (models.Project, error)
	UpdateProjectByAdmin(project models.Project) error
	DeleteProjectByAdmin(project models.Project) error
}

// function create project (with ORM)
func (r *repository) CreateProject(project models.Project) error {
	err := r.db.Create(&project).Error

	return err
}

// function get all projects data in table projects (with ORM)
func (r *repository) FindProject() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Find(&projects).Error

	return projects, err
}

// function select specific data in table project (with ORM), checking by field id
func (r *repository) GetProject(id int) (models.Project, error) {
	var project models.Project
	err := r.db.First(&project, "id = ?", id).Error

	return project, err
}

// function update project by admin if admin will be update (with ORM)
func (r *repository) UpdateProjectByAdmin(project models.Project) error {
	err := r.db.Save(&project).Error

	return err
}

// function delete project by admin if admin will be delete (with ORM)
func (r *repository) DeleteProjectByAdmin(project models.Project) error {
	err := r.db.Delete(&project).Error

	return err
}
