package controllers

import (
	"errors"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var body Body
	// Get the username/email/password
	if err := c.ShouldBindJSON(&body); err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": err})
		return
	}

	err, code := SignupAndStore(c, body)

	if err != nil {
		c.HTML(code, "signup.html", gin.H{"error": err})
	}

	// redirect to the configuration of the account
	c.Redirect(http.StatusFound, "first_connection")
}

func SignupAndStore(c *gin.Context, body Body) (error, int) {
	hash, err := utils.HasPassword(body.Password)

	if err != nil {
		return errors.New("failed to hash password"), http.StatusInternalServerError
	}

	// Create the user
	user := models.User{Role: "member", Email: body.Email, Password: string(hash)}

	// Set a default username
	user.Username = utils.CreateUniqueUsername(body.Email)

	result := initializer.DB.Create(&user)
	if result.Error != nil {
		return errors.New("this user already exist"), http.StatusBadRequest
	}
	//auth the user
	utils.CreateJWT(c, &user)

	return nil, http.StatusOK
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
	utils.CreateJWT(c, &user)
	c.Redirect(http.StatusFound, "/user")
}

func Authorize(c *gin.Context, body Body) (error, int) {
	// Look up requested user
	var user models.User
	initializer.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		return errors.New("user do not exist"), http.StatusBadRequest
	}

	// Compare sent in password with saved password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return errors.New("invalid password"), http.StatusUnauthorized
	}

	// set a jwt
	utils.CreateJWT(c, &user)

	return nil, http.StatusOK
}

func Login(c *gin.Context) {
	// Get the username/email/password
	var body Body

	if err := c.Bind(&body); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": err})
		return
	}

	message, errorCode := Authorize(c, body)

	if errorCode != http.StatusOK {
		c.HTML(errorCode, "login.html", gin.H{"error": message})
	}

	c.Redirect(http.StatusFound, "/user")
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
