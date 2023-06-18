package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InterceptorResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := c.Writer.Status()
		data := c.GetString("data")
		links := c.MustGet("links").([]Link)
		switch statusCode {
		case http.StatusOK: // Get List ~> 200
			c.JSON(statusCode, ResponseGetListData{
				Items:      data,
				Links:      links,
				TotalItems: 1,
				TotalPages: 1,
			})
			break
		case http.StatusCreated: // Created ~> 201
			c.JSON(statusCode, ResponsePostData{
				Items: data,
				Links: links,
			})
			break
		case http.StatusAccepted: // Created ~> 201
			link := c.MustGet("link").(Link)
			c.JSON(statusCode, Link{
				Href:   link.Href,
				Rel:    link.Rel,
				Method: link.Method,
			})
			break
		case http.StatusUnauthorized: // Unauthorized ~> 401
			unauthorizedError := c.MustGet("validationError").(ResponseUnauthorizedData)
			c.JSON(statusCode, ResponseUnauthorizedData{
				Error:            unauthorizedError.Error,
				ErrorDescription: unauthorizedError.ErrorDescription,
			})
			break
		default: // Bad Request Validation ~> 400 | Not Found ~> 401 | Unprocessable Entity ~> 422
			validationError := c.MustGet("validationError").(ResponseErrorData)
			c.JSON(statusCode, ResponseErrorData{
				Name:    validationError.Name,
				Message: validationError.Message,
				DebugId: validationError.DebugId,
				Details: validationError.Details,
				Links:   links,
			})
			break
		}
	}
}
