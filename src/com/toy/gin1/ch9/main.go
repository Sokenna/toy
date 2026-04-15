package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *AppError) Error() string {
	return a.Message
}

var (
	ErrNotFound = &AppError{
		Status:  404,
		Code:    "NOT_FOUND",
		Message: "resource not found",
	}
	ErrUnauthorized = &AppError{
		Status:  401,
		Code:    "UNAUTHORIZED",
		Message: "authentication required",
	}
	ErrBadRequest = &AppError{
		Status:  402,
		Code:    "BAD_REQUEST",
		Message: "invalid request",
	}
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err
		var appErr *AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Status, gin.H{
				"success": false,
				"error":   gin.H{"code": appErr.Code, "message": appErr.Message},
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "INTERNAL", "message": "an unexpected error occurred"}})
		}
	}
}
func VersionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := c.GetHeader("Accept_Version")
		if version == "" {
			version = "v1"
		}
		c.Set("api_version", version)
		c.Next()
	}
}
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v1", "users": []gin.H{}})
		})
	}
	v2 := r.Group("/api/v2")
	{
		v2.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v2", "data": []gin.H{}, "meta": gin.H{"Total": 0}})
		})
	}
	r.Use(VersionMiddleware())
	r.GET("/api/users", func(c *gin.Context) {
		version := c.GetString("api_version")
		switch version {
		case "v2":
			c.JSON(http.StatusOK, gin.H{"version": "v2", "data": []gin.H{}})
		default:
			c.JSON(http.StatusOK, gin.H{"version": "v1", "data": []gin.H{}})
		}
	})
	r.Use(ErrorHandler())
	r.GET("/api/items/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "0" {
			_ = c.Error(ErrNotFound)
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id}})
	})
	r.Run("localhost:8000")
}
