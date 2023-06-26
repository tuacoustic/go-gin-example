package youtubetranscripts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuacoustic/go-gin-example/databases"
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/constants/errorConstants"
	"github.com/tuacoustic/go-gin-example/utils/validate"
)

func Create(c *gin.Context) {
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, errorConstants.DatabaseConnectionError().ErrorName, errorConstants.DatabaseConnectionError().Message, []app.Detail{})
		return
	}
	repo := YoutubeTransriptsRepo(db)

	// Body
	var dataInput YoutubeTranscriptsDto
	if err := c.ShouldBind(&dataInput); err != nil {
		validate.HandleValidationErrors(c, err)
		return
	}

	func(ytbRepo YoutubeTransriptsRepoIF) {
		userInfo := c.MustGet("user").(entities.User)
		item, err := ytbRepo.Create(dataInput, userInfo)
		if err != nil {
			var details []app.Detail
			details = append(details, errorConstants.DatabaseHandlerError(err))
			appG.ErrorResponse(http.StatusBadRequest, errorConstants.YtbTransError().ErrorName, errorConstants.YtbTransError().Message, details)
			return
		}
		appG.Response(http.StatusCreated, []interface{}{item}, app.Pagination{})
	}(repo)
}

func GetDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, errorConstants.DatabaseConnectionError().ErrorName, errorConstants.DatabaseConnectionError().Message, []app.Detail{})
		return
	}
	repo := YoutubeTransriptsRepo(db)
	// Query String
	var queryParams YtbTransQueryParamsDto

	// Bind query parameters to the GetUsersDto struct
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		validate.HandleValidationErrors(c, err)
		return
	}
	func(ytbRepo YoutubeTransriptsRepoIF) {
		item, err := ytbRepo.GetDetail(queryParams)
		if err != nil {
			var details []app.Detail
			details = append(details, errorConstants.DatabaseHandlerError(err))
			appG.ErrorResponse(http.StatusBadRequest, errorConstants.YtbTransError().ErrorName, errorConstants.YtbTransError().Message, details)
			return
		}
		appG.Response(http.StatusOK, []interface{}{item}, app.Pagination{})
	}(repo)
}
