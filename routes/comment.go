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

func CommentRoutes(r *gin.RouterGroup, db *sql.DB) {

	sf := r.Group("/comment")
	sf.Use(middlewares.JWTMiddleware())

	sf.POST("/create", func(c *gin.Context) {

		var comment models.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(comment); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.CommentCreate(c, db, comment)

	})

	sf.GET("/getAll", func(c *gin.Context) {
		controllers.CommentGetAll(c, db)
	})

	sf.POST("/update", func(c *gin.Context) {
		var comment models.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "info": "Invalid input"})
			return
		}

		if err := validate.Struct(comment); err != nil {
			var validationErrors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("Field %s failed on the %s validation", fieldError.Field(), fieldError.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		controllers.CommentUpdate(c, db, comment)
	})

	sf.GET("/getDetail/:id", func(c *gin.Context) {
		controllers.CommentGetByID(c, db)
	})

	sf.DELETE("/delete/:id/:user_id", func(c *gin.Context) {
		controllers.CommentDelete(c, db)
	})
}
