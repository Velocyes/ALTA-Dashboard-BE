package delivery

import (
	"alta-dashboard-be/features/logs"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/consts"
	"alta-dashboard-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LogHandler struct {
	logService logs.LogServiceInterface_
	jwt        middlewares.JWTMiddleware_
}

func New(logService logs.LogServiceInterface_, jwt middlewares.JWTMiddleware_) logs.LogDeliveryInterface_ {
	return &LogHandler{
		logService: logService,
		jwt:        jwt,
	}
}

func (logHandler *LogHandler) AddLog(c echo.Context) error {
	logInput := logs.LogRequest{}
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
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse(consts.LOG_SuccessAddLogData, entityToResponse(logEntity)))
}

func (logHandler *LogHandler) GetLogDataByMenteeId(c echo.Context) error {
	page, limit, err := helper.ExtractPageLimit(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	menteeId, err := helper.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	limit, offset := helper.LimitOffsetConvert(page, limit)
	dataResponse, err := logHandler.logService.GetData(uint(menteeId), limit, offset)
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
