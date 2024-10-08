package common

import "net/http"

type APIResponse struct {
	Message any         `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Status  int         `json:"status"`
}

func GetStatusAcceptedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: http.StatusText(http.StatusOK),
		Status:  http.StatusOK,
		Data:    data,
	}
}

func GetEnvelope(status int, data any) *APIResponse {
	if status >= 400 {
		return &APIResponse{
			Error:   http.StatusText(status),
			Status:  status,
			Message: data,
		}
	}
	return &APIResponse{
		Message: http.StatusText(status),
		Status:  status,
		Data:    data,
	}
}

func GetStatusBadRequestResponse(data any) *APIResponse {
	return &APIResponse{
		Message: http.StatusText(http.StatusBadRequest),
		Status:  http.StatusBadRequest,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetUnauthorizedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: http.StatusText(http.StatusUnauthorized),
		Status:  http.StatusUnauthorized,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetSuccessResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}

func GetInternalServerErrorResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Data:    data,
	}
}

func GetStatusConflictResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusConflict,
		Message: http.StatusText(http.StatusConflict),
		Data:    data,
	}
}
