package validate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/common"
	"github.com/tuacoustic/go-gin-example/utils/constants/errorConstants"
)

var (
	validate = validator.New()
	trans    ut.Translator
)

func Validator(c *gin.Context) {
	c.Set("validator", validate)

	c.Next()
}

func HandleValidationErrors(c *gin.Context, err error) {
	appG := app.Gin{C: c}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors := err.(validator.ValidationErrors)
	var message string
	var details []app.Detail
	for _, e := range errors {
		// var message string
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("The '%s' field is required", common.CamelToSnake(e.Field()))
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		case "email":
			message = fmt.Sprintf("The '%s' must be a valid email, value: %s", common.CamelToSnake(e.Field()), e.Value())
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		case "min":
			message = fmt.Sprintf("The '%s' must be at least %s characters long", common.CamelToSnake(e.Field()), e.Param())
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		case "len":
			message = fmt.Sprintf("The '%s' must be %s characters long", common.CamelToSnake(e.Field()), e.Param())
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		case "gte":
			count := common.CountInterface(e.Value())
			message = fmt.Sprintf("The '%s' must be larger than %d", common.CamelToSnake(e.Field()), count)
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		case "lte":
			count := common.CountInterface(e.Value())
			message = fmt.Sprintf("The '%s' must be smaller than %d", common.CamelToSnake(e.Field()), count)
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		default:
			message = e.Translate(trans)
			details = append(details, errorConstants.CommonDetailError(common.CamelToSnake(e.Field()), e.Value(), errorConstants.CommonBodyLocationError(), message))
		}
	}
	appG.ErrorResponse(http.StatusBadRequest, errorConstants.Validate().ErrorName, errorConstants.Validate().Message, details)
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
