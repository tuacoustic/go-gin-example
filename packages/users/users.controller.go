package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuacoustic/go-gin-example/databases"
	"github.com/tuacoustic/go-gin-example/utils/app"
	"github.com/tuacoustic/go-gin-example/utils/console"
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
	db, err := databases.MysqlConnect()
	if err != nil {
		return
	}
	repo := UsersRepo(db)

	// Body
	var userInput UsersDto
	if err := c.ShouldBind(&userInput); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		validate.HandleValidationErrors(c, err)
		return
	}

	func(userRepo UsersRepoIF) {
		data, err := userRepo.Create(userInput)
		if err != nil {
			return
		}
		console.Info(data)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	func(userRepo UsersRepoIF) {
		items, err := userRepo.GetAll(queryParams)
		if err != nil {
			return
		}
		appG.Response(http.StatusOK, items)
	}(repo)
}
