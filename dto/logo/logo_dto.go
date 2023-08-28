package dtoLogo

type LogoRequestDTO struct {
	Image string `json:"image" form:"image" validate:"required"`
}

type LogoResponseDTO struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
}
