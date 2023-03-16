package response

type JsonResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
