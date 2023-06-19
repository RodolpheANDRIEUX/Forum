package middleware

import (
	"fmt"
	"forum/internal/controllers"
	"forum/internal/initializer"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		controllers.Logout(c)
		c.Redirect(http.StatusUnauthorized, "/")
	}

	// Decode/Validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			controllers.Logout(c)
			c.Redirect(http.StatusUnauthorized, "/")
		}

		// Find the user with token sub id
		var user models.User
		initializer.DB.First(&user, claims["userid"])

		if user.UserID == 0 {
			controllers.Logout(c)
			c.Redirect(http.StatusUnauthorized, "/")
		}
		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		controllers.Logout(c)
		c.Redirect(http.StatusUnauthorized, "/")
	}
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
