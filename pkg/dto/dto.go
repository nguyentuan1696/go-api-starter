package dto

type Pagination[T any] struct {
	Items      []T `json:"items"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
	PageNumber int `json:"page_number "`
	PageSize   int `json:"page_size"`
}

type BaseResponse struct {
	
}


// TODO: Implement BaseRequest
