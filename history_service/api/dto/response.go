package dto

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response[v any] struct {
	BaseResponse
	Result *v `json:"result,omitempty" `
}
