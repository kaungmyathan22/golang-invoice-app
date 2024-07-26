package common

type PaginationParamsRequest struct {
	Page     int
	PageSize int
}

func NewPaginationParamsRequest() *PaginationParamsRequest {
	return &PaginationParamsRequest{
		Page:     1,
		PageSize: 10,
	}
}
