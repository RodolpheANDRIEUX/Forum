package utils

import (
	"errors"
	"forum/internal/initializer"
	"forum/internal/middleware"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
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
	user.UserID = uint(claims["userid"].(float64))

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

	initializer.DB.First(&user, claims["userid"])

	if user.UserID == 0 {
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

// CreateJWT : Create a JWT and set it
func CreateJWT(c *gin.Context, user *models.User) {
	// Generate a jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.UserID,
		"user":   user.Username,
		"email":  user.Email,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24 * 10).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to create token"})
		return
	}
	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*10, "", "", true, true)
}

func AddQuery(link string, queryName string, queryValue string) string {
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	query := u.Query()
	query.Set(queryName, queryValue)
	u.RawQuery = query.Encode()

	return u.String()
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := initializer.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetAllUsersExcept(userID uint) ([]models.User, error) {
	var users []models.User
	err := initializer.DB.Where("user_id != ?", userID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
