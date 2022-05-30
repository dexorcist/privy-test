package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"privy-test/enum"
)

type MultiLangError struct {
	ErrorID int               `json:"error_id"`
	Msg     map[string]string `json:"message"`
}

// MultiCommonInternalError For HTTP Status Internal Server Error
type MultiCommonInternalError MultiLangError

func (c *MultiCommonInternalError) Code() int {
	return c.ErrorID
}

func (c *MultiCommonInternalError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiCommonInternalError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func DefaultMultiInternalError(errorCode enum.ResponseCode) *MultiCommonInternalError {
	return &MultiCommonInternalError{
		ErrorID: errorCode.Int(),
		Msg:     errorCode.StringMap(errorCode),
	}
}

// MultiCommonBadError For HTTP Status Bad Request Error
type MultiCommonBadError MultiLangError

func (c *MultiCommonBadError) Code() int {
	return c.ErrorID
}

func (c *MultiCommonBadError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiCommonBadError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func MultiStringBadError(errorCode enum.ResponseCode) *MultiCommonBadError {
	return &MultiCommonBadError{
		ErrorID: errorCode.Int(),
		Msg:     errorCode.StringMap(errorCode),
	}
}

// MultiCommonNotFoundError For HTTP Status Not Found Error
type MultiCommonNotFoundError MultiLangError

func (c *MultiCommonNotFoundError) Code() int {
	return c.ErrorID
}

func (c *MultiCommonNotFoundError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiCommonNotFoundError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func MultiStringNotFoundError(errorCode enum.ResponseCode) *MultiCommonNotFoundError {
	return &MultiCommonNotFoundError{
		ErrorID: errorCode.Int(),
		Msg:     errorCode.StringMap(errorCode),
	}
}

func IsMultiStringHTTPError(err error, c echo.Context) error {
	switch err.(type) {
	case *MultiCommonBadError:
		return c.JSON(http.StatusBadRequest, err)
	case *MultiCommonNotFoundError:
		return c.JSON(http.StatusNotFound, err)
	case *MultiCommonInternalError:
		return c.JSON(http.StatusInternalServerError, err)
	default:
		return c.JSON(http.StatusInternalServerError, DefaultMultiInternalError(enum.HTTPErrorInternalServerError))
	}
}
