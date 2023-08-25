package dtoResult

// data transfer object result error
type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// data transfer object result success
type SuccessResult struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
