package repositories

import "hollyways/models"

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) error
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(id int) (models.Transaction, error)
	UpdateTransaction(status string, id int) error
	DeleteTransaction(transaction models.Transaction) error
}

func (r *repository) CreateTransaction(transaction models.Transaction) error {
	err := r.db.Create(&transaction).Error

	return err
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Project").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(id int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Project").First(&transaction, "id = ?", id).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, id int) error {
	var transaction models.Transaction
	r.db.First(&transaction, id)

	if status != transaction.Status && status == "success" {
		var project models.Project
		r.db.First(project, transaction.ProjectID)

		project.Donation += transaction.Donation
		r.db.Save(&project)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) error {
	err := r.db.Delete(&transaction).Error

	return err
}
