package common

import "net/http"

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetStatusAcceptedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "success",
		Status:  http.StatusOK,
		Data:    data,
	}
}

func GetStatusBadRequestResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "error",
		Status:  http.StatusBadRequest,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetUnauthorizedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "error",
		Status:  http.StatusUnauthorized,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetSuccessResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func GetNotFoundResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusNotFound,
		Message: "error",
		Data:    data,
	}
}

func GetForbiddenResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusForbidden,
		Message: "error",
		Data:    data,
	}
}
