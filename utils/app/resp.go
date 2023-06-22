package app

import "net/http"

func (g *Gin) Response(httpCode int, data interface{}, pagination Pagination) {
	switch httpCode {
	case http.StatusOK: // Get List ~> 200
		g.C.JSON(httpCode, ResponseGetListData{
			Items: data,
			// Links:      links,
			CurrentPage: pagination.CurrentPage,
			TotalItems:  pagination.TotalItems,
			TotalPages:  pagination.TotalPages,
		})
	case http.StatusCreated: // Created ~> 201
		g.C.JSON(httpCode, ResponsePostData{
			Items: data,
			// Links: links,
		})
		// case http.StatusAccepted: // Created ~> 202
		// 	g.C.JSON(httpCode, Link{
		// 		Href:   link.Href,
		// 		Rel:    link.Rel,
		// 		Method: link.Method,
		// 	})
		// 	break
		// case http.StatusUnauthorized: // Unauthorized ~> 401
		// 	unauthorizedError := c.MustGet("validationError").(ResponseUnauthorizedData)
		// 	g.C.JSON(httpCode, ResponseUnauthorizedData{
		// 		Error:            unauthorizedError.Error,
		// 		ErrorDescription: unauthorizedError.ErrorDescription,
		// 	})
		// 	break
		// default: // Bad Request Validation ~> 400 | Not Found ~> 401 | Unprocessable Entity ~> 422
		// 	validationError := c.MustGet("validationError").(ResponseErrorData)
		// 	g.C.JSON(httpCode, ResponseErrorData{
		// 		Name:    validationError.Name,
		// 		Message: validationError.Message,
		// 		DebugId: validationError.DebugId,
		// 		Details: validationError.Details,
		// 		Links:   links,
		// 	})
		// 	break
	}
}
func (g *Gin) ErrorResponse(httpCode int, name string, message string, details []Detail) {
	switch httpCode {
	default: // Bad Request Validation ~> 400 | Not Found ~> 401 | Unprocessable Entity ~> 422
		g.C.JSON(httpCode, ResponseErrorData{
			Name:    name,
			Message: message,
			Details: details,
		})
	}
}
