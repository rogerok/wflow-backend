package models

type PaginationQuery struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}
