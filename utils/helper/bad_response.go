package helper

import (
	"alta-dashboard-be/utils/consts"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ValidateUserFailedResponse(c echo.Context, err error) (codeStatus int, failedMessage string) {
	if err.Error() == consts.DATABASE_InvalidQueryParameter {
		return http.StatusBadRequest, consts.DATABASE_InvalidQueryParameter
	} else if err.Error() == consts.VALIDATION_InvalidInput {
		return http.StatusBadRequest, consts.VALIDATION_InvalidInput
	} else if err.Error() == consts.SERVER_ForbiddenRequest {
		return http.StatusBadRequest, consts.SERVER_ForbiddenRequest
	} else if err.Error() == consts.USER_EmptyCredentialError {
		return http.StatusBadRequest, consts.USER_EmptyCredentialError
	} else if err.Error() == gorm.ErrRecordNotFound.Error() {
		return http.StatusNotFound, gorm.ErrRecordNotFound.Error()
	} else if err.Error() == consts.USER_EmailAlreadyUsed {
		return http.StatusBadRequest, consts.USER_EmailAlreadyUsed
	}
	return http.StatusInternalServerError, err.Error()
}

func ValidateLogFailedResponse(c echo.Context, err error) (codeStatus int, failedMessage string) {
	if err.Error() == consts.VALIDATION_InvalidInput {
		return http.StatusBadRequest, consts.VALIDATION_InvalidInput
	} else if err.Error() == consts.LOG_InvalidParamStatus {
		return http.StatusBadRequest, consts.LOG_InvalidParamStatus
	} else if err.Error() == consts.LOG_MenteeNotExisted {
		return http.StatusBadRequest, consts.LOG_MenteeNotExisted
	}
	
	return http.StatusInternalServerError, err.Error()
}
