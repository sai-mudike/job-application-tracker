package middelware

import (
	"github.com/gin-gonic/gin"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/handlers"
	"github.com/sai-mudike/job-application-tracker/internals/services"
)

func Authentication(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		handlers.HandleError(context, appErrors.ErrMissingToken)
		return
	}

	userID, err := services.VerifyToken(token)

	if err != nil {
		handlers.HandleError(context, appErrors.ErrInvalidToken)
		return
	}

	context.Set("userID", userID)

	context.Next()

}
