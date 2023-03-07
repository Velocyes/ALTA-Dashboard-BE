package delivery

import (
	"alta-dashboard-be/features/users"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/consts"
	"alta-dashboard-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService users.UserServiceInterface_
	jwt         middlewares.JWTMiddleware_
}

func New(userService users.UserServiceInterface_, jwt middlewares.JWTMiddleware_) users.UserDeliveryInterface_ {
	return &UserHandler{
		userService: userService,
		jwt:         jwt,
	}
}

func (userHandler *UserHandler) Login(c echo.Context) error {
	loginInput := users.UserLogin{}
	err := c.Bind(&loginInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.USER_ErrorBindUserData))
	}

	userEntity, token, err := userHandler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if err.Error() == consts.USER_EmptyCredentialError {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_EmptyCredentialError))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	dataResponse := map[string]any{
		"token": token,
		"data":  entityToResponse(userEntity),
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(consts.USER_LoginSuccess, dataResponse))
}

func (userHandler *UserHandler) Register(c echo.Context) error {
	userInput := users.UserRegister{}
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.USER_ErrorBindUserData))
	}

	inputedUserEntity := registerToEntity(userInput)

	_, loggedInUserRole, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	userEntity, err := userHandler.userService.Create(inputedUserEntity, loggedInUserRole)
	if err != nil {
		if err.Error() == consts.SERVER_ForbiddenRequest {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
		} else if err.Error() == consts.USER_EmptyCredentialError {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_EmptyCredentialError))
		} else if err.Error() == consts.USER_EmailAlreadyUsed {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_EmailAlreadyUsed))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse(consts.USER_RegisterSuccess, entityToResponse(userEntity)))
}

func (userHandler *UserHandler) GetAllUser(c echo.Context) error {
	page, limit, err := helper.ExtractPageLimit(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse, err := userHandler.userService.GetAll(helper.LimitOffsetConvert(page, limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	userEntities, exist := dataResponse["data"]
	if !exist {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.SERVER_InternalServerError))
	}
	dataResponse["data"] = listEntityToResponse(userEntities.([]users.UserEntity))
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(consts.USER_SuccessReadUserData, dataResponse))
}

func (userHandler *UserHandler) GetUserData(c echo.Context) error {
	userId, errExtract := helper.ExtractIDParam(c)
	if errExtract != nil {
		return c.JSON(http.StatusBadRequest, errExtract)
	}

	loggedInUserId, loggedInUserRole, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	userEntity, err := userHandler.userService.GetData(loggedInUserId, uint(userId), loggedInUserRole)
	if err != nil {
		if err.Error() == consts.SERVER_ForbiddenRequest {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse(consts.USER_UserNotFound))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.SERVER_InternalServerError))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(consts.USER_SuccessReadUserData, entityToResponse(userEntity)))
}

func (userHandler *UserHandler) UpdateAccount(c echo.Context) error {
	userId, errExtract := helper.ExtractIDParam(c)
	if errExtract != nil {
		return c.JSON(http.StatusBadRequest, errExtract)
	}

	loggedInUserId, loggedInUserRole, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	userInput := users.UserRequest{}
	err = c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_ErrorBindUserData))
	}
	userEntity := requestToEntity(userInput)

	userEntity, err = userHandler.userService.ModifyData(loggedInUserId, uint(userId), loggedInUserRole, userEntity)
	if err != nil {
		if err.Error() == consts.SERVER_ForbiddenRequest {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse(consts.USER_UserNotFound))
		} else if err.Error() == consts.USER_FailedUpdate {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.USER_FailedUpdate))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.SERVER_InternalServerError))
	}

	return c.JSON(http.StatusNoContent, helper.SuccessWithDataResponse(consts.USER_SuccessUpdateUserData, entityToResponse(userEntity)))
}

func (userHandler *UserHandler) RemoveAccount(c echo.Context) error {
	userId, errExtract := helper.ExtractIDParam(c)
	if errExtract != nil {
		return c.JSON(http.StatusBadRequest, errExtract)
	}

	loggedInUserId, loggedInUserRole, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	err = userHandler.userService.Remove(loggedInUserId, uint(userId), loggedInUserRole)
	if err != nil {
		if err.Error() == consts.SERVER_ForbiddenRequest {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(consts.SERVER_ForbiddenRequest))
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse(consts.USER_UserNotFound))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(consts.USER_FailedDelete))
	}

	return c.JSON(http.StatusNoContent, helper.SuccessResponse(consts.USER_SuccessDelete))
}
