package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tuacoustic/go-gin-example/docs"
	"github.com/tuacoustic/go-gin-example/middlewares"
	"github.com/tuacoustic/go-gin-example/packages/auth"
	"github.com/tuacoustic/go-gin-example/packages/users"
	youtubetranscripts "github.com/tuacoustic/go-gin-example/packages/youtubeTranscripts"
	"github.com/tuacoustic/go-gin-example/utils/validate"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": []map[string]string{
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
	commonUri := "/api/v1"
	usersGroup := r.Group(commonUri)
	authGroup := r.Group(commonUri)
	youtubeTranscriptsGroup := r.Group(commonUri)
	// Authentication Middleware
	authGroup.Use(middlewares.AuthMiddleware())
	youtubeTranscriptsGroup.Use(middlewares.AuthMiddleware())
	{
		// User Routes
		usersGroup.POST("/users/register", users.Create)
		usersGroup.GET("/users/get-all", users.GetAll)
		usersGroup.PUT("/users/:id/update", users.Update)
		usersGroup.DELETE("/users/:id/soft-delete", users.SoftDelete)
	}
	{
		// Auth Routes
		usersGroup.POST("/auth/login", auth.Login)
		authGroup.GET("/auth/profile", auth.Profile)
	}
	{
		// Youtube Transcript Routes
		youtubeTranscriptsGroup.POST("/youtube/transcripts/create", youtubetranscripts.Create)
		youtubeTranscriptsGroup.GET("/youtube/transcripts/get-detail", youtubetranscripts.GetDetail)
	}

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
