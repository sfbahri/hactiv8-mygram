package controllers

import (
	"database/sql"
	"log"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context, db *sql.DB, comment models.Comment) {

	comment.CreatedAt = time.Now()

	var photoId int
	queryPhoto := `SELECT id FROM photo WHERE user_id = $1`
	err := db.QueryRow(queryPhoto, comment.UserID).Scan(&photoId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Photo not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get photo URL"})
		return
	}

	queryComment := `INSERT INTO comment (user_id, photo_id, message, created_at) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(queryComment, comment.UserID, photoId, comment.Message, comment.CreatedAt).Scan(&comment.ID)
	if err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"info":   "Created Comment Success",
	})

}

func CommentGetAll(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, user_id, photo_id, message, created_at FROM comment")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get data"})
		return
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			log.Printf("Scan error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error scanning data"})
			return
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Error with rows"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   comments,
	})
}

func CommentUpdate(c *gin.Context, db *sql.DB, comment models.Comment) {

	id := comment.ID

	query := `UPDATE comment SET message = $1, updated_at = $2 WHERE id = $3`
	comment.UpdatedAt = time.Now()

	// Execute the update statement
	result, err := db.Exec(query, comment.Message, comment.UpdatedAt, id)
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
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Update Comment Success",
	})
}

func CommentGetByID(c *gin.Context, db *sql.DB) {

	id := c.Param("id")

	query := `SELECT id, user_id, photo_id, message, created_at FROM comment WHERE id = $1`
	var comment models.Comment

	err := db.QueryRow(query, id).Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Comment not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   comment,
	})
}

func CommentDelete(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	usr_id := c.Param("user_id")

	query := `SELECT id FROM comment WHERE id = $1 and user_id = $2`
	var commentID uint

	err := db.QueryRow(query, id, usr_id).Scan(&commentID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "info": "Comment not found"})
			return
		}
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to get comment"})
		return
	}

	deleteQuery := `DELETE FROM comment WHERE id = $1`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "info": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "Comment deleted successfully",
	})
}
