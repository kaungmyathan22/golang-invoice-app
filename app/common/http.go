package common

type PaginationParamsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PaginationMeta struct {
	TotalItems   int64 `json:"totalItems"`
	TotalPages   int   `json:"totalPages"`
	Page         int   `json:"page"`
	PageSize     int   `json:"pageSize"`
	NextPage     *int  `json:"nextPage"`
	PreviousPage *int  `json:"previousPage"`
}

func (pagination *PaginationParamsRequest) SetDefaultPaginationValues() {
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
}

func (pagination *PaginationParamsRequest) GetMeta(totalRecords int64) *PaginationMeta {
	totalPages := int((totalRecords + int64(pagination.PageSize) - 1) / int64(pagination.PageSize))
	var nextPage, previousPage *int

	// Determine next and previous page numbers
	if pagination.Page < totalPages {
		next := pagination.Page + 1
		nextPage = &next
	}
	if pagination.Page > 1 {
		prev := pagination.Page - 1
		previousPage = &prev
	}

	return &PaginationMeta{
		TotalPages:   totalPages,
		TotalItems:   totalRecords,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		Page:         pagination.Page,
		PageSize:     pagination.PageSize,
	}
}

func NewPaginationParamsRequest() *PaginationParamsRequest {
	return &PaginationParamsRequest{
		Page:     1,
		PageSize: 10,
	}
}

type PaginationResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Data     any `json:"data"`
}
