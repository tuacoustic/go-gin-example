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
	}
}
func (g *Gin) ErrorResponse(httpCode int, name string, message string, details []Detail) {
	switch httpCode {
	default: // Bad Request Validation ~> 400 | Not Found ~> 404 | Unprocessable Entity ~> 422
		g.C.JSON(httpCode, ResponseErrorData{
			Name:    name,
			Message: message,
			Details: details,
		})
	}
}
func (g *Gin) UnauthorizedResponse(httpCode int, errorName string, errorDesc string) {
	g.C.JSON(httpCode, ResponseUnauthorizedData{
		Error:            errorName,
		ErrorDescription: errorDesc,
	})
}
