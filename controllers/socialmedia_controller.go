package controllers

import (
	"database/sql"
	"log"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SocialMediaCreate(c *gin.Context, db *sql.DB, socialmedia models.SocialMedia) {

	socialmedia.CreatedAt = time.Now()

	querySocialMedia := `INSERT INTO socialmedia (user_id, name, social_media_url, created_at) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(querySocialMedia, socialmedia.UserID, socialmedia.Name, socialmedia.SocialMediaURL, socialmedia.CreatedAt).Scan(&socialmedia.ID)
	if err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"info":   "Created Social Media Success",
	})

}

func SocialMediaGetAll(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, social_media_url, user_id,created_at FROM socialmedia")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get data"})
		return
	}
	defer rows.Close()

	var socialmedia []models.SocialMedia
	for rows.Next() {
		var socmed models.SocialMedia
		err := rows.Scan(&socmed.ID, &socmed.Name, &socmed.SocialMediaURL, &socmed.UserID, &socmed.CreatedAt)
		if err != nil {
			log.Printf("Scan error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error scanning data"})
			return
		}
		socialmedia = append(socialmedia, socmed)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error with rows"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   socialmedia,
	})
}

func SocialMediaUpdate(c *gin.Context, db *sql.DB, socmed models.SocialMedia) {

	id := socmed.ID

	query := `UPDATE socialmedia SET name = $1,social_media_url = $2,user_id = $3, updated_at = $4 WHERE id = $5`
	socmed.UpdatedAt = time.Now()

	// Execute the update statement
	result, err := db.Exec(query, socmed.Name, socmed.SocialMediaURL, socmed.UserID, socmed.UpdatedAt, id)
	if err != nil {
		log.Printf("Update error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to update social media"})
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error checking affected rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to check affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Social media not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Update Social Media Success",
	})
}

func SocialMediaGetByID(c *gin.Context, db *sql.DB) {

	id := c.Param("id")

	query := `SELECT id, name, social_media_url, user_id, created_at FROM socialmedia WHERE id = $1`
	var socmed models.SocialMedia
	err := db.QueryRow(query, id).Scan(&socmed.ID, &socmed.Name, &socmed.SocialMediaURL, &socmed.UserID, &socmed.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Social Media not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get Social Media"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   socmed,
	})
}

func SocialMediaDelete(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	query := `SELECT id FROM socialmedia WHERE id = $1`
	var socmedID uint

	err := db.QueryRow(query, id).Scan(&socmedID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Social Media not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get socialmedia"})
		return
	}

	deleteQuery := `DELETE FROM socialmedia WHERE id = $1`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to delete socialmedia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Social Media deleted successfully",
	})
}
