package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tuacoustic/go-gin-example/docs"
	"github.com/tuacoustic/go-gin-example/packages/users"
	"github.com/tuacoustic/go-gin-example/utils/validate"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": []map[string]string{
			{
				"version_number": "v1.0.1",
			},
		},
	})
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Home page
	r.GET("/", HomePage)

	// Interceptor
	r.Use(validate.Validator)
	apiv1 := r.Group("/api/v1")
	{
		// User Routes
		apiv1.POST("/users/register", users.Create)
		apiv1.GET("/users/get-all", users.GetAll)
	}

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
