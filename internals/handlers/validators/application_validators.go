package validators

import (
	"net/url"
	"strings"
	"time"

	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func ValidateCreateApplication(req models.Application) error {

	// Trim strings
	req.Company_name = strings.TrimSpace(req.Company_name)
	req.Job_title = strings.TrimSpace(req.Job_title)
	req.Job_url = strings.TrimSpace(req.Job_url)
	req.Location = strings.TrimSpace(req.Location)
	req.Employment_type = strings.TrimSpace(req.Employment_type)
	req.Status = strings.TrimSpace(req.Status)
	req.Notes = strings.TrimSpace(req.Notes)

	// Company name
	if req.Company_name == "" {
		return appErrors.ErrInvalidApplicationData
	}

	if len(req.Company_name) > 100 {
		return appErrors.ErrInvalidApplicationData
	}

	// Job title
	if req.Job_title == "" {
		return appErrors.ErrInvalidApplicationData
	}

	if len(req.Job_title) > 150 {
		return appErrors.ErrInvalidApplicationData
	}

	// Job URL
	if req.Job_url == "" {
		return appErrors.ErrInvalidApplicationData
	}

	u, err := url.Parse(req.Job_url)
	if err != nil || u.Host == "" {
		return appErrors.ErrInvalidJob_url
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return appErrors.ErrInvalidJob_url
	}

	// Location
	if req.Location == "" {
		return appErrors.ErrInvalidApplicationData
	}

	if len(req.Location) > 100 {
		return appErrors.ErrInvalidApplicationData
	}

	// Employment type
	if !isValidEmployment_type(req.Employment_type) {
		return appErrors.ErrInvalidEmployment_type
	}

	// Salary minimum
	if req.Salary_min != 0 && req.Salary_min < 0 {
		return appErrors.ErrInvalidSalaryRange
	}

	// Salary maximum
	if req.Salary_max != 0 && req.Salary_max < 0 {
		return appErrors.ErrInvalidSalaryRange
	}

	// Salary range
	if req.Salary_min != 0 &&
		req.Salary_max != 0 &&
		req.Salary_min > req.Salary_max {

		return appErrors.ErrInvalidSalaryRange
	}

	// Status
	if !isValidStatus(req.Status) {
		return appErrors.ErrInvalidStatus
	}

	// Applied date
	if req.Applied_at.IsZero() {

		if req.Applied_at.After(time.Now()) {
			return appErrors.ErrInvalidApplied_at
		}
	}

	// Notes
	if len(req.Notes) > 1000 {
		return appErrors.ErrInvalidApplicationData
	}

	return nil
}

func isValidEmployment_type(value string) bool {

	switch strings.ToLower(value) {
	case "full_time",
		"part_time",
		"contract",
		"internship",
		"temporary":
		return true
	default:
		return false
	}

}

func isValidStatus(value string) bool {
	switch strings.ToLower(value) {
	case "applied",
		"interview",
		"offer",
		"rejected",
		"withdrawn",
		"accepted":
		return true
	default:
		return false
	}
}
