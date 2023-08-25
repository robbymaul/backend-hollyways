package dtoAds

type AdsRequestDTO struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}

type AdsResponseDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type AdsUpdateRequestDTO struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}
