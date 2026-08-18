package errors

import "errors"

var (
	// 400 Bad Request
	ErrInvalidRequest         = errors.New("invalid request")
	ErrInvalidApplicationData = errors.New("invalid application data")
	ErrInvalidSalaryRange     = errors.New("invalid salary range")
	ErrInvalidStatus          = errors.New("invalid application status")
	ErrInvalidJob_url         = errors.New("invalid job url")
	ErrInvalidEmployment_type = errors.New("invalid employment type")
	ErrInvalidApplied_at      = errors.New("invalid Applied at")

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

	// 409 Conflict
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrDuplicateApplication = errors.New("application already exists")

	// 500 Internal Server Error
	ErrInternal = errors.New("internal server error")
)
