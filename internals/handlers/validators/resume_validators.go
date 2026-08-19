package validators

import (
	"strings"

	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func ValidateCreateResume(req *models.Resume) error {
	req.Name = strings.TrimSpace(req.Name)
	req.File_path = strings.TrimSpace(req.File_path)

	if req.Name == "" {
		return appErrors.ErrInvalidResumeData
	}
	if len(req.Name) > 255 {
		return appErrors.ErrInvalidResumeData

	}

	if req.File_path == "" {
		return appErrors.ErrInvalidResumeData

	}
	return nil

}
