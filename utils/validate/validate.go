package validate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/common"
	"github.com/tuacoustic/go-gin-example/utils/constants/validateConstants"
)

var (
	validate = validator.New()
	trans    ut.Translator
)

func Validator(c *gin.Context) {
	c.Set("validator", validate)

	c.Next()
}

func detailValidation(field string, value string, location string, description string) app.DetailValidation {
	return app.DetailValidation{
		Field:       field,
		Value:       value,
		Location:    location,
		Description: description,
	}
}

func HandleValidationErrors(c *gin.Context, err error) {
	appG := app.Gin{C: c}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors := err.(validator.ValidationErrors)
	// validationErrors := make(map[string]string)
	var message string
	var details app.DetailValidation
	for _, e := range errors {
		// var message string
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("The '%s' field is required", common.CamelToSnake(e.Field()))
			details = detailValidation(common.CamelToSnake(e.Field()), "", e.Field(), "Field is required")
		case "email":
			message = fmt.Sprintf("The '%s' must be a valid email, value: %s", common.CamelToSnake(e.Field()), e.Value())
			details = detailValidation(common.CamelToSnake(e.Field()), "", e.Field(), "Field is required")
		case "min":
			message = fmt.Sprintf("The '%s' must be at least %s characters long", common.CamelToSnake(e.Field()), e.Param())
			details = detailValidation(common.CamelToSnake(e.Field()), "", e.Field(), "Field is required")
		case "len":
			message = fmt.Sprintf("The '%s' must be %s characters long", common.CamelToSnake(e.Field()), e.Param())
			details = detailValidation(common.CamelToSnake(e.Field()), "", e.Field(), "Field is required")
		default:
			message = e.Translate(trans)
			details = detailValidation(common.CamelToSnake(e.Field()), "", e.Field(), "Field is required")
		}
		// validationErrors[common.CamelToSnake(e.Field())] = message
	}
	// c.JSON(http.StatusBadRequest, gin.H{"errors": message})
	// c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
	// details := app.DetailValidation{
	// 	Field:
	// }
	appG.ErrorResponse(http.StatusBadRequest, validateConstants.Validate().ErrorName, message, details)
}

// {
//     "name": "INVALID_REQUEST",
//     "message": "Request is not well-formed, syntactically incorrect, or violates schema.",
//     "debug_id": "b6b9a374802ea",
//     "details": [
//         {
//             "field": "/intent",
//             "value": "SHIPPING",
//             "location": "body",
//             "issue": "INVALID_PARAMETER_VALUE",
//             "description": "The value of a field is invalid."
//         }
//     ],
//     "links": [
//         {
//             "href": "https://developer.paypal.com/docs/api/orders/v2/#error-INVALID_PARAMETER_VALUE",
//             "rel": "information_link",
//             "encType": "application/json"
//         }
//     ]
// }
