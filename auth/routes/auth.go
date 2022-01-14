package routes

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/rojasleon/reserve-micro/auth/internal"
	"github.com/rojasleon/reserve-micro/auth/models"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func InitAuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", signup)
		auth.POST("/signin", signin)
		auth.GET("/currentuser", currentUser)
	}
}

func signup(c *gin.Context) {
	user := models.User{}

	// Populate user data
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong params",
		})
		return
	}

	validate = validator.New()

	// Validate incoming credentials
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	existingUser := models.User{}

	err := models.DB.First(&existingUser, "email = ?", user.Email).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	models.DB.Create(&user)

	// Generate JWT
	token, err := internal.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	c.Header("Authorization", token)

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func signin(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal error",
		})
		return
	}

	validate = validator.New()

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	existingUser := models.User{}

	err := models.DB.First(&existingUser, "email = ?", user.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exist",
		})
		return
	}

	// Generate JWT
	token, err := internal.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	c.Header("Authorization", token)

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func currentUser(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"currentUser": "null",
		})
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	claims := internal.VerifyJWT(token)

	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"currentUser": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"currentUser": claims,
	})
}
