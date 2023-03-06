package delivery

import (
	"alta-dashboard-be/features/logs"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/consts"
	"alta-dashboard-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LogHandler struct {
	logService logs.LogServiceInterface
}

func New(logService logs.LogServiceInterface) *LogHandler {
	return &LogHandler{
		logService: logService,
	}
}

func (logHandler *LogHandler) AddLog(c echo.Context) error {
	logInput := LogRequest{}
	err := c.Bind(&logInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.LOG_ErrorBindLogData))
	}

	inputedLogEntity := requestToEntity(logInput)

	loggedInUserId, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	logEntity, err := logHandler.logService.Create(inputedLogEntity, loggedInUserId)
	if err != nil {
		if err.Error() == consts.SERVER_ForbiddenRequest {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse(consts.LOG_SuccessAddLogData, entityToResponse(logEntity)))
}

func (logHandler *LogHandler) GetLogDataByMenteeId(c echo.Context) error {
	page, limit, pageParam, limitParam, menteeIdParam := -1, -1, c.QueryParam("page"), c.QueryParam("limit"), c.Param("mentee_id")
	searchedMenteeId, err := strconv.Atoi(menteeIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.ECHO_InvalidParam))
	}
	if pageParam != "" {
		castedPageParam, errCasting := strconv.Atoi(pageParam)
		if errCasting != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.ECHO_InvalidParam))
		}
		page = castedPageParam
	}
	if limitParam != "" {
		castedLimitParam, errCasting := strconv.Atoi(limitParam)
		if errCasting != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.ECHO_InvalidParam))
		}
		limit = castedLimitParam
	}

	limit, offset := helper.LimitOffsetConvert(page, limit)
	dataResponse, err := logHandler.logService.GetData(uint(searchedMenteeId), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	logEntities, exist := dataResponse["data"]
	if !exist {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.SERVER_InternalServerError))
	}
	dataResponse["data"] = listEntityToResponse(logEntities.([]logs.LogEntity))
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(consts.USER_SuccessReadUserData, dataResponse))
}

// func (logHandler *LogHandler) UpdateLog(c echo.Context) error {
// 	logId, errCasting := strconv.Atoi(c.Param("log_id"))
// 	if errCasting != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.ECHO_InvalidParam))
// 	}

// 	loggedInUserId, loggedInUserRole, err := middlewares.ExtractToken(c)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
// 	}

// 	userInput := UserRequest{}
// 	err = c.Bind(&userInput)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_ErrorBindUserData))
// 	}
// 	userEntity := requestToEntity(userInput)

// 	userEntity, err = userHandler.userService.ModifyData(loggedInUserId, uint(userId), loggedInUserRole, userEntity)
// 	if err != nil {
// 		if err.Error() == consts.SERVER_ForbiddenRequest{
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
// 		} else if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, helper.FailedResponse(consts.USER_UserNotFound))
// 		} else if err.Error() == consts.USER_FailedUpdate{
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_FailedUpdate))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.SERVER_InternalServerError))
// 	}

// 	return c.JSON(http.StatusNoContent, helper.SuccessWithDataResponse(consts.USER_SuccessUpdateUserData, entityToResponse(userEntity)))
// }

// func (userHandler *UserHandler) RemoveAccount(c echo.Context) error {
// 	userId, errCasting := strconv.Atoi(c.Param("id"))
// 	if errCasting != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.ECHO_InvalidParam))
// 	}

// 	loggedInUserId, loggedInUserRole, err := middlewares.ExtractToken(c)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
// 	}

// 	err = userHandler.userService.Remove(loggedInUserId, uint(userId), loggedInUserRole)
// 	if err != nil {
// 		if err.Error() == consts.SERVER_ForbiddenRequest{
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
// 		} else if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, helper.FailedResponse(consts.USER_UserNotFound))
// 		} else if err.Error() == consts.USER_FailedDelete{
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_FailedUpdate))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.USER_FailedDelete))
// 	}

// 	return c.JSON(http.StatusNoContent, helper.SuccessResponse(consts.USER_SuccessDelete))
// }
