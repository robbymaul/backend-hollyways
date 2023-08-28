package models

import "gorm.io/gorm"

// models structure database table transactions
type Transaction struct {
	gorm.Model
	UserID    int     `json:"user_id" gorm:"type: int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User    `json:"user" gorm:"foreignKey: UserID"`
	ProjectID int     `json:"project_id" gorm:"type: int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Project   Project `json:"project" gorm:"foreignKey: ProjectID"`
	Donation  int     `json:"donation" gorm:"type: int"`
	Status    string  `json:"status" gorm:"type: varchar(255);default:pending"`
}

// models response if table joining relation schema
type TransactionResponse struct {
	User     User    `json:"user"`
	Project  Project `json:"project"`
	Donation int     `json:"donation"`
}

// function for hande not create new table transactions
func (TransactionResponse) TableName() string {
	return `transactions`
}
