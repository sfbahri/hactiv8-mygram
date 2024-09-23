package controllers

import (
	"database/sql"
	"log"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PhotoCreate(c *gin.Context, db *sql.DB, photo models.Photo) {

	photo.CreatedAt = time.Now()

	query := `INSERT INTO photo (title, caption, photo_url, user_id, created_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, photo.Title, photo.Caption, photo.PhotoURL, photo.UserID, photo.CreatedAt).Scan(&photo.ID)
	if err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"info":   "Created Photo Success",
		"data":   photo,
	})

}

func PhotoUpdate(c *gin.Context, db *sql.DB, photo models.Photo) {

	id := photo.ID

	query := `UPDATE photo SET title = $1, caption = $2, photo_url = $3, user_id = $4, updated_at = $5 WHERE id = $6`
	photo.UpdatedAt = time.Now() // Set the updated_at field

	// Execute the update statement
	result, err := db.Exec(query, photo.Title, photo.Caption, photo.PhotoURL, photo.UserID, photo.UpdatedAt, id)
	if err != nil {
		log.Printf("Update error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to update photo"})
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
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Update Photo Success",
		"data":   photo,
	})
}

func PhotoGetAll(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, title, caption, photo_url, user_id, created_at FROM photo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get data"})
		return
	}
	defer rows.Close()

	var photos []models.Photo
	for rows.Next() {
		var photo models.Photo
		err := rows.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt)
		if err != nil {
			log.Printf("Scan error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error scanning data"})
			return
		}
		photos = append(photos, photo)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error with rows"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   photos,
	})
}

func PhotoGetByID(c *gin.Context, db *sql.DB) {

	id := c.Param("id")

	query := `SELECT id, title, caption, photo_url, user_id, created_at FROM photo WHERE id = $1`
	var photo models.Photo

	err := db.QueryRow(query, id).Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Photo not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   photo,
	})
}

func PhotoDelete(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	// First, check if the photo exists
	query := `SELECT id FROM photo WHERE id = $1`
	var photoID uint

	err := db.QueryRow(query, id).Scan(&photoID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Photo not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get photo"})
		return
	}

	// Photo exists, proceed to delete
	deleteQuery := `DELETE FROM photo WHERE id = $1`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Photo deleted successfully",
	})
}
