package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuacoustic/go-gin-example/databases"
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/constants/errorConstants"
	"github.com/tuacoustic/go-gin-example/utils/validate"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		appG.ErrorResponse(http.StatusInternalServerError, errorConstants.DatabaseConnectionError().ErrorName, errorConstants.DatabaseConnectionError().Message, []app.Detail{})
		return
	}
	repo := AuthRepo(db)

	// Body
	var authInput AuthDto
	if err := c.ShouldBind(&authInput); err != nil {
		validate.HandleValidationErrors(c, err)
		return
	}

	func(authRepo AuthRepoIF) {
		item, err := authRepo.Login(authInput)
		if err != nil {
			var details []app.Detail
			details = append(details, errorConstants.DatabaseHandlerError(err))
			appG.ErrorResponse(http.StatusBadRequest, errorConstants.AuthError().ErrorName, errorConstants.AuthError().Message, details)
			return
		}
		appG.Response(http.StatusOK, []interface{}{item}, app.Pagination{})
	}(repo)
}

func Profile(c *gin.Context) {
	appG := app.Gin{C: c}
	item := c.MustGet("user").(entities.User)
	appG.Response(http.StatusOK, []interface{}{item}, app.Pagination{})
}
