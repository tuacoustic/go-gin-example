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

func AuthError() CommonStruct {
	return CommonStruct{
		ErrorName: "AUTH_ERROR",
		Message:   "Auth error.",
	}
}

func YtbTransError() CommonStruct {
	return CommonStruct{
		ErrorName: "YOUTUBE_TRANSCRIPT_ERROR",
		Message:   "Youtube transcript error.",
	}
}

func DatabaseConnectionError() CommonStruct {
	return CommonStruct{
		ErrorName: "DATABASE_ERROR",
		Message:   "No connection could be made because the target machine actively refused it.",
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

func DatabaseHandlerError(err error) app.Detail {
	errorMessage := err.Error()
	var resultResp app.Detail
	// Define the regular expression pattern
	pattern := `Duplicate entry '([^']+)' for key '([^']+)'`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)
	// Find submatches using the regular expression
	matches := regex.FindStringSubmatch(errorMessage)

	// Extract the desired parts
	if len(matches) >= 3 {
		resultResp = app.Detail{
			Field:       matches[1],
			Value:       matches[2],
			Location:    CommonBodyLocationError(),
			Description: regex.FindString(errorMessage),
		}
	} else {
		resultResp = app.Detail{
			Field:       "",
			Value:       "",
			Location:    CommonBodyLocationError(),
			Description: errorMessage,
		}
	}
	return resultResp
}

func NotNullJsonError() app.Detail {
	notNullJsonResp := app.Detail{
		Field:       "",
		Value:       "",
		Location:    CommonBodyLocationError(),
		Description: commonConstants.NotNullBodyMsg(),
	}
	return notNullJsonResp
}
