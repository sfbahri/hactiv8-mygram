package routes

import (
	"database/sql"
	"fmt"
	"mygram/controllers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func UserRoutes(r *gin.RouterGroup, db *sql.DB) {
	r.POST("/user/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(user); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		if controllers.IsDuplicateUser(db, user.Username, user.Email) {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}

		controllers.Register(c, db, user)
	})

	r.POST("/user/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Username == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"info": "Pastikan username dan password tidak kosong!"})
			return
		}

		controllers.Auth(c, db, user)

	})
}
