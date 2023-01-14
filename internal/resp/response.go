package resp

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}
type JSONBadReqResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong parameter"`
	Data    interface{} `json:"data" `
}
type JSONIntServerErrReqResult struct {
	Code    int         `json:"code" example:"404"`
	Message string      `json:"message" example:"This was not found"`
	Data    interface{} `json:"data" `
}

func SuccessResponse(c echo.Context, code int, data interface{}) error {
	c.JSON(http.StatusCreated, JSONSuccessResult{
		Code:    code,
		Data:    data,
		Message: "Success",
	})
	return nil
}

func FailResponse(c echo.Context, respCode int, message string) error {
	if respCode == http.StatusInternalServerError {
		c.JSON(respCode, JSONIntServerErrReqResult{
			Code:    respCode,
			Data:    nil,
			Message: message,
		})
		return nil
	}
	c.JSON(respCode, JSONBadReqResult{
		Code:    respCode,
		Data:    nil,
		Message: message,
	})
	return nil
}
