package delivery

import (
	"alta-dashboard-be/features/mentee"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	service mentee.MenteeService_
}

// Create implements mentee.MenteeDelivery_
func (u *MenteeDelivery) Create(c echo.Context) error {
	//extract userID from jwt
	userID, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("invalid or expired jwt"))
	}

	//bind payload
	var req mentee.MenteeRequest
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
	err = u.service.Create(int(userID), core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("created"))
}

// Delete implements mentee.MenteeDelivery_
func (u *MenteeDelivery) Delete(c echo.Context) error {
	//extract userID from jwt
	userID, _, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("invalid or expired jwt"))
	}
	//extract id from param
	id, err := helper.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	err = u.service.Delete(int(userID), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("deleted"))
}

// GetAll implements mentee.MenteeDelivery_
func (u *MenteeDelivery) GetAll(c echo.Context) error {
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
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("", convertToResponseList(data)))
}

// GetOne implements mentee.MenteeDelivery_
func (u *MenteeDelivery) GetOne(c echo.Context) error {
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
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("", convertToResponse(&data)))
}

// Update implements mentee.MenteeDelivery_
func (u *MenteeDelivery) Update(c echo.Context) error {
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
	var req mentee.MenteeRequest
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

func New(service mentee.MenteeService_) mentee.MenteeDelivery_ {
	return &MenteeDelivery{
		service: service,
	}
}
