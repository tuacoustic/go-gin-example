package users

import (
	"math"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tuacoustic/go-gin-example/databases"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/constants/errorConstants"
	"github.com/tuacoustic/go-gin-example/utils/validate"
)

// CreateUsersHandler handles the POST /api/v1/users/register endpoint.
// ShowAccount godoc
// @Summary      Register a new user
// @Description  register a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        email   path      string  true  "email"
// @Param        phone   path      string  true  "phone"
// @Param        avatar   path      string  true  "avatar"
// @Param        password   path      string  true  "password"
// @Success      200  {object}  entities.User
// @Failure      400  {object}  resp.ResponseErrorData
// @Router       /api/v1/users/register [post]
func Create(c *gin.Context) {
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		return
	}
	repo := UsersRepo(db)

	// Body
	var userInput UsersDto
	if err := c.ShouldBind(&userInput); err != nil {
		validate.HandleValidationErrors(c, err)
		return
	}

	func(userRepo UsersRepoIF) {
		item, err := userRepo.Create(userInput)
		if err != nil {
			var details []app.Detail
			details = append(details, errorConstants.DuplicateError(err))
			appG.ErrorResponse(http.StatusBadRequest, errorConstants.UserError().ErrorName, errorConstants.UserError().Message, details)
			return
		}
		appG.Response(http.StatusCreated, item, app.Pagination{})
	}(repo)
}

func GetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		return
	}
	repo := UsersRepo(db)

	// Query String
	var queryParams GetUsersDto

	// Bind query parameters to the GetUsersDto struct
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		validate.HandleValidationErrors(c, err)
		return
	}

	func(userRepo UsersRepoIF) {
		items, count, err := userRepo.GetAll(queryParams)
		if err != nil {
			return
		}
		totalPages := math.Round(float64(count) / float64(queryParams.Limit))
		pagination := app.Pagination{
			CurrentPage: queryParams.Page,
			TotalItems:  count,
			TotalPages:  int(totalPages),
		}
		appG.Response(http.StatusOK, items, pagination)
	}(repo)
}
func Update(c *gin.Context) {
	userId := c.Param("id")
	appG := app.Gin{C: c}
	db, err := databases.MysqlConnect()
	if err != nil {
		return
	}
	repo := UsersRepo(db)
	// Body
	var userInput UpdateUserDto
	if err := c.ShouldBind(&userInput); err != nil || reflect.DeepEqual(userInput, UpdateUserDto{}) {
		var details []app.Detail
		if err != nil {
			validate.HandleValidationErrors(c, err)
			return
		}
		details = append(details, errorConstants.NotNullJsonError())
		appG.ErrorResponse(http.StatusBadRequest, errorConstants.UserError().ErrorName, errorConstants.UserError().Message, details)
		return
	}
	func(userRepo UsersRepoIF) {
		item, err := userRepo.Update(userId, userInput)
		if err != nil {
			var details []app.Detail
			details = append(details, errorConstants.DuplicateError(err))
			appG.ErrorResponse(http.StatusBadRequest, errorConstants.UserError().ErrorName, errorConstants.UserError().Message, details)
			return
		}
		appG.Response(http.StatusCreated, item, app.Pagination{})
	}(repo)
}
