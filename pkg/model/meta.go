package model

type Meta struct {
	Page    uint `json:"page"`
	PerPage uint `json:"per_page"`
	Total   uint `json:"total"`
}
