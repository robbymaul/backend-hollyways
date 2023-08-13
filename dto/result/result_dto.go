package dtoResult

type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResult struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
