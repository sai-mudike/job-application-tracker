package models

type ApplicationQuery struct {
	Page           int
	Limit          int
	Status         string
	EmploymentType string
	SortBy         string
	OrderBy        string
}
