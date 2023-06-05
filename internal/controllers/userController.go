package controllers

import (
	"forum/internal/database"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	// Get the username/email/password
	var body struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": err})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to hash password"})
		return
	}
	// Create the user
	user := models.User{Role: "member", Username: body.Username, Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "This user already exist"})
		return
	}
	// Respond
	c.HTML(http.StatusOK, "user.html", gin.H{"username": body.Username})
}
func Login(c *gin.Context) {
	// Get the username/email/password
	var body struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": err})
		return
	}
	// Look up requested user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "User do not exist"})
		return
	}
	// Compare sent in password with saved password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid password"})
		return
	}
	// Generate a jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 10).Unix(),
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
	c.Redirect(http.StatusFound, "/user")
}
func User(c *gin.Context) {
	var user models.User
	id, _ := c.Get("user")
	database.DB.First(&user, "id = ?", id)
	c.HTML(http.StatusOK, "user.html", gin.H{"username": "user"})
}
func Logout(c *gin.Context) {
	// Delete the cookie
	c.SetCookie("Authorization", "", -1, "", "", true, true)
	c.Redirect(http.StatusFound, "/")
}
