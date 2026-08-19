package validators

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func ApplQueryVali(context *gin.Context, query *models.ApplicationQuery, handleError func(*gin.Context, error)) error {

	allowedSorts := map[string]bool{
		"created_at":   true,
		"updated_at":   true,
		"applied_at":   true,
		"company_name": true,
		"salary_min":   true,
		"salary_max":   true,
	}
	allowedEmployment := map[string]bool{
		"full_time":  true,
		"part_time":  true,
		"contract":   true,
		"internship": true,
		"temporary":  true,
	}

	allowedStatus := map[string]bool{
		"applied":   true,
		"interview": true,
		"offer":     true,
		"rejected":  true,
		"withdrawn": true,
		"accepted":  true,
	}
	if page := context.Query("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil || p < 1 {
			return appErrors.ErrInvalidPagination
		}

		query.Page = p
	}
	if limit := context.Query("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil || l < 1 || l > 100 {
			return appErrors.ErrInvalidPagination
		}
		query.Limit = l

	}

	if !allowedSorts[query.SortBy] {
		return appErrors.ErrInvalidSort
	}
	if query.OrderBy != "asc" && query.OrderBy != "desc" {
		return appErrors.ErrInvalidSortOrder
	}
	if query.OrderBy == "asc" {
		query.OrderBy = "ASC"
	}

	if query.OrderBy == "desc" {
		query.OrderBy = "DESC"
	}
	if query.EmploymentType != "" {

		if !allowedEmployment[query.EmploymentType] {

			return appErrors.ErrInvalidEmployment_type
		}
	}
	if query.Status != "" {
		fmt.Println(query.Status)

		if !allowedStatus[query.Status] {

			return appErrors.ErrInvalidStatus
		}
	}
	if query.EmploymentType == "" || query.Status == "" {
		return nil
	}
	return nil
}
