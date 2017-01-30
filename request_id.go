package requestid

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func RequestIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestId := c.Request().Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String() //"aaaa" //uuid.NewV4().String()
		}
		c.Set("RequestId", requestId)
		c.Response().Header().Set("X-Request-Id", requestId)
		c.Request().Header.Set("X-Request-Id", requestId)
		return next(c)
	}
}
