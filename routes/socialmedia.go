package routes

import (
	"database/sql"
	"fmt"
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func SocialMediaRoutes(r *gin.RouterGroup, db *sql.DB) {

	sf := r.Group("/socialmedia")
	sf.Use(middlewares.JWTMiddleware())

	sf.POST("/create", func(c *gin.Context) {

		var socialmedia models.SocialMedia
		if err := c.ShouldBindJSON(&socialmedia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(socialmedia); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.SocialMediaCreate(c, db, socialmedia)

	})

	sf.GET("/getAll", func(c *gin.Context) {
		controllers.SocialMediaGetAll(c, db)
	})

	sf.POST("/update", func(c *gin.Context) {
		var socmed models.SocialMedia
		if err := c.ShouldBindJSON(&socmed); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "info": "Invalid input"})
			return
		}

		if err := validate.Struct(socmed); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.SocialMediaUpdate(c, db, socmed)
	})

	sf.GET("/getDetail/:id", func(c *gin.Context) {
		controllers.SocialMediaGetByID(c, db)
	})

	sf.DELETE("/delete/:id/:user_id", func(c *gin.Context) {
		controllers.SocialMediaDelete(c, db)
	})
}
