package models

// Pager is
type Pager struct {
	//当前页数，0基
	Page int `json:"page"`
	//每页的大小
	PageSize int `json:"pageSize"`
	//总的条目数
	TotalItems int `json:"totalItems"`
	//总的页数
	TotalPages int `json:"totalPages"`
	//实体的数组
	Data interface{} `json:"data"`
}
