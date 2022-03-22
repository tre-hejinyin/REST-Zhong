package errorhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"server/middleware/logger"
	"server/middleware/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			var APIErrors []APIError
			var apiErr *APIError
			var code int
			var errorType ErrorType
			var errorMessage string
			for _, err := range c.Errors {
				switch err.Type {
				case gin.ErrorTypePublic:
					var pgError *pgconn.PgError
					if errors.As(err.Err, &pgError) {
						errorType = ErrDatabase
						errorMessage = MsgErrDatabase
						if pgError.Code == GetPgErrorType(ErrUniqueViolation) {
							errorType = ErrUniqueViolation
							errorMessage = pgError.Detail
						}
					} else if errors.Is(err.Err, sql.ErrNoRows) {
						errorType = ErrNotFound
						errorMessage = sql.ErrNoRows.Error()
					} else if errors.Is(err, gorm.ErrRecordNotFound) {
						errorType = ErrNotFound
						errorMessage = sql.ErrNoRows.Error()
					} else {
						errorType = ErrInternal
						errorMessage = MsgErrInternal
					}
					apiErr = NewAPIError(errorType, errorMessage)
					logger.Errorf("error_type: %s, error_message: %s, detail: %s", apiErr.ErrorType, apiErr.ErrorMessage, err.Error())
					APIErrors = append(APIErrors, *apiErr)
				case gin.ErrorTypeBind:
					errorType = ErrInvalidRequest
					var validationErrors validator.ValidationErrors
					if errors.As(err.Err, &validationErrors) {
						for _, e := range validationErrors {
							errorMessage = ValidationErrorToText(e)
							apiErr = NewAPIError(errorType, errorMessage)
							logger.Warnf("error_type: %s, error_message: %s", apiErr.ErrorType, apiErr.ErrorMessage)
							APIErrors = append(APIErrors, *apiErr)
						}
					} else {
						apiErr = NewAPIError(errorType, err.Error())
						logger.Warnf("error_type: %s, error_message: %s", apiErr.ErrorType, apiErr.ErrorMessage)
						APIErrors = append(APIErrors, *apiErr)
					}
				default:
					var numErr *strconv.NumError
					if errors.As(err.Err, &numErr) {
						errorType = ErrBadParams
						errorMessage = strings.ReplaceAll(MsgErrNumError, "{num}", numErr.Num)
						logger.Warnf("error_type: %s, error_message: %s", errorType, errorMessage)
					} else {
						errorType = ErrInternal
						errorMessage = MsgErrInternal
						logger.Errorf("error_type: %s, error_message: %s, detail: %S", errorType, errorMessage, err.Error())
					}
					apiErr = NewAPIError(errorType, errorMessage)
					APIErrors = append(APIErrors, *apiErr)
				}
			}
			code = GetHTTPStatus(errorType)
			if c.Writer.Status() != http.StatusOK {
				code = c.Writer.Status()
			}
			c.AbortWithStatusJSON(code, response.Body{
				Code: code,
				Msg:  fmt.Sprintf("%+v", APIErrors),
			})
		}
	}
}

func ValidationErrorToText(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("'%s' is required", field)
	case "max":
		return fmt.Sprintf("'%s'cannot be longer than %s", field, e.Param())
	case "min":
		return fmt.Sprintf("'%s'must be longer than %s", field, e.Param())
	case "email":
		return "invalid email format"
	case "len":
		return fmt.Sprintf("'%s' must be %s characters long", field, e.Param())
	}
	return fmt.Sprintf("'%s' is not valid", field)
}
