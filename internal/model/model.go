package model

type ApiResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors any           `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
