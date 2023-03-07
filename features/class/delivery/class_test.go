package delivery

import (
	"alta-dashboard-be/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var BASE_URL = "http://localhost:8081/classes"

func TestDelete(t *testing.T) {
	table := DeleteTable()

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock jwt
			jwtMid := new(mocks.JWTMiddleware_)
			jwtMid.On("ExtractToken", mock.Anything).Return(uint(1), "admin", nil)
			//mock service
			service := new(mocks.ClassService_)
			service.On("Delete", mock.Anything, mock.Anything).Return(nil)

			//create new echo context
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%v", BASE_URL, v.Input.ID), nil)
			rec := httptest.NewRecorder()

			//register endpoint
			e := echo.New()
			delivery := New(service, jwtMid)
			e.DELETE("classes/:id", delivery.Delete)

			//execute delivery
			e.ServeHTTP(rec, req)

			//test
			if v.Output.IsError {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode error
				assert.GreaterOrEqual(t, rec.Result().StatusCode, 400, data)
			} else {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode ok
				assert.LessOrEqual(t, rec.Result().StatusCode, 300, data)
			}
		})
	}
}
func TestUpdate(t *testing.T) {
	table := UpdateTable()

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock jwt
			jwtMid := new(mocks.JWTMiddleware_)
			jwtMid.On("ExtractToken", mock.Anything).Return(uint(1), "admin", nil)
			//mock service
			service := new(mocks.ClassService_)
			service.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)

			//create new echo context
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%v", BASE_URL, v.Input.ID), bytes.NewBuffer(v.Input.Class))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			//register endpoint
			e := echo.New()
			delivery := New(service, jwtMid)
			e.PUT("classes/:id", delivery.Update)

			//execute delivery
			e.ServeHTTP(rec, req)

			//test
			if v.Output.IsError {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode error
				assert.GreaterOrEqual(t, rec.Result().StatusCode, 400, data)
			} else {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode ok
				assert.LessOrEqual(t, rec.Result().StatusCode, 300, data)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	table := CreateTable()

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock jwt
			jwtMid := new(mocks.JWTMiddleware_)
			jwtMid.On("ExtractToken", mock.Anything).Return(uint(1), "admin", nil)
			//mock service
			service := new(mocks.ClassService_)
			service.On("Create", mock.Anything, mock.Anything).Return(nil)

			//create new echo context
			req := httptest.NewRequest(http.MethodPost, BASE_URL, bytes.NewBuffer(v.Input.Class))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			//execute delivery
			delivery := New(service, jwtMid)
			delivery.Create(ctx)

			//test
			if v.Output.IsError {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode error
				assert.GreaterOrEqual(t, rec.Result().StatusCode, 400, data)
			} else {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode ok
				assert.LessOrEqual(t, rec.Result().StatusCode, 300, data)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	table := GetOneTable()

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock jwt
			jwtMid := new(mocks.JWTMiddleware_)
			jwtMid.On("ExtractToken", mock.Anything).Return(1, "admin", nil)
			//mock service
			service := new(mocks.ClassService_)
			service.On("GetOne", mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//create new echo handler
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%v", BASE_URL, v.Input.ID), nil)
			rec := httptest.NewRecorder()

			//register endpoint
			e := echo.New()
			delivery := New(service, jwtMid)
			e.GET("classes/:id", delivery.GetOne)

			//execute delivery
			e.ServeHTTP(rec, req)

			//test
			if v.Output.IsError {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode error
				assert.GreaterOrEqual(t, rec.Result().StatusCode, 400, data)
			} else {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode ok
				assert.LessOrEqual(t, rec.Result().StatusCode, 300, data)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	table := GetAllTable()

	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock jwt
			jwtMid := new(mocks.JWTMiddleware_)
			jwtMid.On("ExtractToken", mock.Anything).Return(1, "admin", nil)
			//mock service
			service := new(mocks.ClassService_)
			service.On("GetAll", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//create new echo context
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?page=%v&limit=%v", BASE_URL, v.Input.Page, v.Input.Limit), nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			//execute delivery
			delivery := New(service, jwtMid)
			delivery.GetAll(ctx)

			//test
			if v.Output.IsError {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode error
				assert.GreaterOrEqual(t, rec.Result().StatusCode, 400, data)
			} else {
				//check payloads
				var data map[string]interface{} = make(map[string]interface{})
				err := json.NewDecoder(rec.Result().Body).Decode(&data)
				assert.Nil(t, err, data)
				//check statusCode ok
				assert.LessOrEqual(t, rec.Result().StatusCode, 300, data)
			}
		})
	}
}
