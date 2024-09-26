package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
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

	// Serve the request
	e.ServeHTTP(w, r) // Serve the HTTP request using Echo
}

// Entry point for the Vercel function
func main() {
	// The main function is not typically needed in Vercel functions
	// as the Handler function is used directly
}
