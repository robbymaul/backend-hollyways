package dtoAuth

type RegisterRequestDTO struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
}
