package response

// APIResponse универсальная структура ответа
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`  // Данные ответа
	Message string      `json:"message"`         // Сообщение о результате
	Error   interface{} `json:"error,omitempty"` // Детали ошибки, если есть
}

// SuccessResponse создаёт успешный ответ
func SuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
		Message: message,
		Error:   nil,
	}
}

// ErrorResponse создаёт ответ с ошибкой
func ErrorResponse(err interface{}, message string) APIResponse {
	return APIResponse{
		Success: false,
		Data:    nil,
		Message: message,
		Error:   err,
	}
}
