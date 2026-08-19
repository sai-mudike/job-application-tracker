package handlers

import (
	"errors"
	"fmt"
	"net/http"

	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {

	switch {

	// 400 Bad Request
	case errors.Is(err, appErrors.ErrInvalidRequest):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: fmt.Sprintf("Invalid request %v", err.Error()),
		})

	case errors.Is(err, appErrors.ErrInvalidApplicationData):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_APPLICATION_DATA",
			Message: "Invalid application data",
		})

	case errors.Is(err, appErrors.ErrInvalidSalaryRange):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SALARY_RANGE",
			Message: "Minimum salary cannot be greater than maximum salary",
		})

	case errors.Is(err, appErrors.ErrInvalidStatus):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_STATUS",
			Message: "Invalid application status",
		})

	case errors.Is(err, appErrors.ErrInvalidJob_url):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_JOB_URL",
			Message: "Invalid Job url "})

	case errors.Is(err, appErrors.ErrInvalidEmployment_type):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_EMPLOYMNET_TYPE",
			Message: "Invalid employment type"})

	case errors.Is(err, appErrors.ErrInvalidApplied_at):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_APPLIED_AT",
			Message: "Applied at Cannot be after the apllication creating data"})
	case errors.Is(err, appErrors.ErrInvalidPagination):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_PAGINATION",
			Message: "Valid page number is required"})
	case errors.Is(err, appErrors.ErrInvalidSortOrder):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SORTING_ORDER",
			Message: "sorting can be either 'asc' or 'desc'"})
	case errors.Is(err, appErrors.ErrInvalidSort):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "INVALID_SORTING",
			Message: "invalid sorting data"})

	// 401 Unauthorized
	case errors.Is(err, appErrors.ErrMissingToken):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "MISSING_TOKEN",
			Message: "Authentication token is required",
		})

	case errors.Is(err, appErrors.ErrInvalidToken):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "INVALID_TOKEN",
			Message: "Invalid authentication token",
		})

	case errors.Is(err, appErrors.ErrTokenExpired):
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "TOKEN_EXPIRED",
			Message: "Authentication token has expired",
		})

	case errors.Is(err, appErrors.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    "INVALID_CREDENTIALS",
			Message: "Invalid email or password",
		})

	// 403 Forbidden
	case errors.Is(err, appErrors.ErrForbidden):
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    "FORBIDDEN",
			Message: "You do not have permission to perform this action",
		})

	// 404 Not Found
	case errors.Is(err, appErrors.ErrUserNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "USER_NOT_FOUND",
			Message: "User not found",
		})

	case errors.Is(err, appErrors.ErrApplicationNotFound):
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    "APPLICATION_NOT_FOUND",
			Message: "Application not found",
		})

	// 409 Conflict
	case errors.Is(err, appErrors.ErrEmailAlreadyExists):
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Code:    "EMAIL_ALREADY_EXISTS",
			Message: "An account with this email already exists",
		})

	case errors.Is(err, appErrors.ErrDuplicateApplication):
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Code:    "DUPLICATE_APPLICATION",
			Message: "You already have this application tracked",
		})

	// 500 Internal Server Error
	default:
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: fmt.Sprintf("An unexpected error occurred %v", err.Error()),
		})
		return
	}
}
