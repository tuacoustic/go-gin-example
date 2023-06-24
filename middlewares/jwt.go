package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tuacoustic/go-gin-example/databases"
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/channel"
	"github.com/tuacoustic/go-gin-example/utils/constants/commonConstants"
	"github.com/tuacoustic/go-gin-example/utils/constants/errorConstants"
	tablename "github.com/tuacoustic/go-gin-example/utils/constants/tableName"
	"github.com/tuacoustic/go-gin-example/utils/setting"
)

func GenerateToken(uuid uuid.UUID) (string, error) {
	payload := jwt.MapClaims{
		"sub": uuid,
		"exp": commonConstants.JwtConstants().Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(string(setting.AppSetting.JwtSecret)))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
		}
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var userData entities.User
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		userUUID := claims["sub"]
		db, err := databases.MysqlConnect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorConstants.DatabaseConnectionError().Message})
			c.Abort()
			return
		}
		done := make(chan bool)
		go func(ch chan<- bool) {
			defer close(ch)
			if err = db.Debug().Table(tablename.TableName().Users).Where("uuid = ?", userUUID).First(&userData).Error; err != nil {
				ch <- false
				return
			}
			ch <- true
		}(done)
		if channel.OK(done) {
			c.Set("user", userData)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
		c.Abort()
	}
}
