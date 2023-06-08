package utils

import (
	"errors"
	"forum/internal/initializer"
	"forum/internal/middleware"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HasPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	return hash, err
}

func ParseUser(c *gin.Context) (models.User, error) {
	var user = models.User{}

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		return user, err
	}

	claims, err := middleware.ParseToken(tokenString)
	if err != nil {
		return user, err
	}

	user.Username = claims["user"].(string)
	user.Email = claims["email"].(string)
	user.Role = claims["role"].(string)

	return user, nil
}

func GetUSer(c *gin.Context) (models.User, error) {
	var user = models.User{}

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		return user, err
	}

	claims, err := middleware.ParseToken(tokenString)
	if err != nil {
		return user, err
	}

	initializer.DB.First(&user, claims["id"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return user, errors.New("401")
	}

	return user, nil
}

func CreateUniqueUsername(email string) string {
	emailUsername := strings.Split(email, "@")[0]

	// generate a random number
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(99999-10001) + 10000

	// merge both
	return emailUsername + strconv.Itoa(number)
}
