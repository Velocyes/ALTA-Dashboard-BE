package helper

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LimitOffsetConvert(page, limit int) (int, int) {
	offset := -1
	if limit > 0 {
		offset = (page - 1) * limit
	}
	return limit, offset
}

func ExtractIDParam(c echo.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func ExtractPageLimit(c echo.Context) (page int, limit int, err error) {
	pageStr := c.QueryParam("page")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, errors.New("invalid page parameter")
	}
	limitStr := c.QueryParam("limit")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, errors.New("invalid limit parameter")
	}
	return page, limit, nil
}
