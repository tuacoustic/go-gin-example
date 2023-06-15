package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tuacoustic/go-gin-example/docs"
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
	// apiv1 := r.Group("/api/v1")
	// {
	// 	apiv1.GET("/", HomePage)
	// }

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
