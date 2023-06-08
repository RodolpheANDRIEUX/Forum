package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// todo : reiriger vers username si unername n'existe pas mais le mail oui
// todo : gmail logic new auth system

func Signup(c *gin.Context) {
	// Get the username/email/password
	var body struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": err})
		return
	}
	// Hash the password
	hash, err := utils.HasPassword(body.Password)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	user := models.User{Role: "member", Email: body.Email, Password: string(hash)}

	// Set a default username
	user.Username = utils.CreateUniqueUsername(body.Email)

	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "This user already exist"})
		return
	}
	//auth the user
	CreateJWT(c, &user)

	// redirect to the configuration of the account
	c.Redirect(http.StatusFound, "first_connection")
}

func SendUsername(c *gin.Context) {
	user, err := utils.GetUSer(c)

	if err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": err})
		Logout(c)
		return
	}

	newUsername := c.PostForm("username")
	user.Username = newUsername

	result := initializer.DB.Save(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "This username already exist"})
		return
	}
	CreateJWT(c, &user)
	c.Redirect(http.StatusFound, "/user")
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
	initializer.DB.First(&user, "email = ?", body.Email)
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

	// set a jwt
	CreateJWT(c, &user)

	c.Redirect(http.StatusFound, "/user")
}

// CreateJWT : Create a JWT and set it
func CreateJWT(c *gin.Context, user *models.User) {
	// Generate a jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"user":  user.Username,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24 * 10).Unix(),
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

func User(c *gin.Context) {
	user, err := utils.ParseUser(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.HTML(http.StatusOK, "user.html", gin.H{"user": user})
}

func Logout(c *gin.Context) {
	// Delete the cookie
	c.SetCookie("Authorization", "", -1, "", "", true, true)
	c.Redirect(http.StatusFound, "/")
}
