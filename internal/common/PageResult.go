package common

import "go-micro-demo/internal/common/response"

type PageResult struct {
	List     []*response.UserList `json:"list"`
	Total    uint64               `json:"total"`
	Page     uint64               `json:"page"`
	PageSize uint64               `json:"pageSize"`
}
