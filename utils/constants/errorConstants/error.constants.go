package errorConstants

import (
	"regexp"

	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/constants/commonConstants"
)

func Validate() CommonStruct {
	return CommonStruct{
		ErrorName: "INVALID_REQUEST",
		Message:   "Request is not well-formed, syntactically incorrect, or violates schema.",
	}
}

func UserError() CommonStruct {
	return CommonStruct{
		ErrorName: "USER_ERROR",
		Message:   "User error.",
	}
}

func CommonBodyLocationError() string {
	return "body"
}

func CommonDetailError(field string, value interface{}, location string, description string) app.Detail {
	return app.Detail{
		Field:       field,
		Value:       value,
		Location:    location,
		Description: description,
	}
}

func DuplicateError(err error) app.Detail {
	errorMessage := err.Error()
	var duplicateResp app.Detail
	// Define the regular expression pattern
	pattern := `Duplicate entry '([^']+)' for key '([^']+)'`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Find submatches using the regular expression
	matches := regex.FindStringSubmatch(errorMessage)

	// Extract the desired parts
	if len(matches) >= 3 {
		duplicateResp = app.Detail{
			Field:       matches[1],
			Value:       matches[2],
			Location:    CommonBodyLocationError(),
			Description: regex.FindString(errorMessage),
		}
	}
	return duplicateResp
}

func NotNullJsonError() app.Detail {
	var notNullJsonResp app.Detail
	notNullJsonResp = app.Detail{
		Field:       "",
		Value:       "",
		Location:    CommonBodyLocationError(),
		Description: commonConstants.NotNullBodyMsg(),
	}
	return notNullJsonResp
}
