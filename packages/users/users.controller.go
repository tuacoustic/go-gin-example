package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuacoustic/go-gin-example/databases"
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
// @Failure      400  {object}  badRequestResponse
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if err := validate.Validate().Struct(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	func(userRepo UsersRepoIF) {
		data, err := userRepo.Create(userInput)
		if err != nil {
			return
		}
		console.Info(data)
		return
	}(repo)
}

func GetAll(c *gin.Context) {
	db, err := databases.MysqlConnect()
	if err != nil {
		return
	}
	repo := UsersRepo(db)

	// Query String
	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams)
	func(userRepo UsersRepoIF) {
		data, err := userRepo.GetAll(queryParams)
		if err != nil {
			return
		}
		console.Info(data)
		return
	}(repo)
}

// swagger:response badRequestResponse
type badRequestResponse struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	} `json:"body"`
}
