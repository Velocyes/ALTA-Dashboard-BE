package delivery

import (
	"alta-dashboard-be/features/class"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	service class.ClassService_
}

// Create implements class.ClassDelivery_
func (u *ClassDelivery) Create(c echo.Context) error {
	//get user id from jwt
	id, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("invalid or expired token"))
	}
	//bind payload
	var req class.ClassRequest
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("unknown payload"))
	}

	//convert to core
	core, err := convertToCore(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	//create
	err = u.service.Create(int(id), core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("class created"))
}

// Delete implements class.ClassDelivery_
func (u *ClassDelivery) Delete(c echo.Context) error {
	//get user id from jwt
	userID, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("invalid or expired token"))
	}
	//get id from param
	id, err := helper.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	//delete
	err = u.service.Delete(int(userID), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("class created"))
}

// GetAll implements class.ClassDelivery_
func (u *ClassDelivery) GetAll(c echo.Context) error {
	//get limit and page from query param
	page, limit, err := helper.ExtractPageLimit(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	//get all data
	data, err := u.service.GetAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	//convert to response list
	d, err := convertToResponseList(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, d)
}

// GetOne implements class.ClassDelivery_
func (u *ClassDelivery) GetOne(c echo.Context) error {
	//get id from param
	id, err := helper.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	//get one data
	data, err := u.service.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	res, err := convertToResponse(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, res)
}

// Update implements class.ClassDelivery_
func (u *ClassDelivery) Update(c echo.Context) error {
	//get user id from jwt
	userID, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("invalid or expired token"))
	}
	//get id from param
	id, err := helper.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	//bind payload
	var req class.ClassRequest
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("unknown payload"))
	}

	//convert to core
	core, err := convertToCore(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	//update
	err = u.service.Update(int(userID), id, core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse(fmt.Sprintf("class with id %d updated", id)))
}

func New(service class.ClassService_) class.ClassDelivery_ {
	return &ClassDelivery{
		service: service,
	}
}
