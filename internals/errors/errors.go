package errors

import (
	"errors"

	"github.com/lib/pq"
)

var PQErr *pq.Error

var (
	// 400 Bad Request
	ErrInvalidRequest         = errors.New("invalid request")
	ErrInvalidApplicationData = errors.New("invalid application data")
	ErrInvalidSalaryRange     = errors.New("invalid salary range")
	ErrInvalidStatus          = errors.New("invalid application status")
	ErrInvalidJob_url         = errors.New("invalid job url")
	ErrInvalidEmployment_type = errors.New("invalid employment type")
	ErrInvalidApplied_at      = errors.New("invalid Applied at")
	ErrInvalidPagination      = errors.New("invalid page number")
	ErrInvalidSort            = errors.New("invalid sorting value")
	ErrInvalidSortOrder       = errors.New("invalid sorting order")
	ErrInvalidResumeData      = errors.New("invalid Resume data")
	ErrInvalidResumeFile_path = errors.New("invalid resume file path")

	// 401 Unauthorized
	ErrMissingToken       = errors.New("missing token")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidCredentials = errors.New("invalid credentials")

	// 403 Forbidden
	ErrForbidden = errors.New("forbidden")

	// 404 Not Found
	ErrUserNotFound        = errors.New("user not found")
	ErrApplicationNotFound = errors.New("application not found")
	ErrResumeNotFound      = errors.New("Resume not found")

	// 409 Conflict
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrDuplicateApplication = errors.New("application already exists")
	ErrResumeAlreadyExists  = errors.New("resume already exists")

	// 500 Internal Server Error
	ErrInternal = errors.New("internal server error")
)
