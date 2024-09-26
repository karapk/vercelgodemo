package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	e := echo.New()

	// Middleware for recovery
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					if httpError, ok := err.(echo.HTTPError); ok {
						c.JSON(httpError.Code, map[string]string{
							"message": httpError.Error(),
						})
					} else {
						message := fmt.Sprintf("%s", err)
						c.JSON(http.StatusInternalServerError, map[string]string{
							"message": message,
						})
					}
				}
			}()
			return next(c)
		}
	}))

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	})

	e.GET("/hello", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "name not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"data": fmt.Sprintf("Hello %s!", name),
		})
	})

	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": map[string]string{
				"id": id,
			},
		})
	})

	e.GET("/long/long/long/path/*test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": map[string]string{
				"url": c.Path(),
			},
		})
	})

	// Start the server
	return e.Start(":8080") // You may change the port as needed
}
