package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute handlers

		// Check if there are any errors attached to the context
		if len(c.Errors) > 0 {
			var errors []string
			for _, e := range c.Errors {
				errors = append(errors, e.Error())
			}

			// If the status code is still 200 (OK), but we have errors, force an error code.
			// Usually the handler might set the code, but if not, we default to 500 or 400.
			// However, ideally handlers set the status code.
			// If status is 200, assume Internal Server Error for unhandled errors
			if c.Writer.Status() == http.StatusOK {
				c.Status(http.StatusInternalServerError)
			}

			c.JSON(c.Writer.Status(), ErrorResponse{
				Success: false,
				Errors:  errors,
			})
		}
	}
}


