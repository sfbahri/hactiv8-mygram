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

func PhotoRoutes(r *gin.RouterGroup, db *sql.DB) {

	sf := r.Group("/photo")
	sf.Use(middlewares.JWTMiddleware())

	sf.GET("/getAll", func(c *gin.Context) {
		controllers.PhotoGetAll(c, db)
	})

	sf.POST("/create", func(c *gin.Context) {

		var photo models.Photo
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(photo); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.PhotoCreate(c, db, photo)

	})

	sf.GET("/getDetail/:id", func(c *gin.Context) {
		controllers.PhotoGetByID(c, db)
	})

	sf.POST("/update", func(c *gin.Context) {
		var photo models.Photo
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "info": "Invalid input"})
			return
		}

		if err := validate.Struct(photo); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.PhotoUpdate(c, db, photo)
	})

	sf.DELETE("/delete/:id/:user_id", func(c *gin.Context) {
		controllers.PhotoDelete(c, db)
	})
}
