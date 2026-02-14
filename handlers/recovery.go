package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[PANIC RECOVERED] %v\n%s", r, debug.Stack())
					_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"data": nil,
						"error": map[string]interface{}{
							"message": "an unexpected internal error occurred",
							"code":    "INTERNAL_ERROR",
						},
					})
				}
			}()
			return next(c)
		}
	}
}
