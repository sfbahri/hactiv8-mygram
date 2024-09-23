package controllers

import (
	"database/sql"
	"log"
	"mygram/lib"
	"mygram/models"
	"mygram/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func IsDuplicateUser(db *sql.DB, username, email string) bool {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2`
	err := db.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return false // If there's an error in the query, treat it as no duplicate
	}
	return count > 0
}

func Register(c *gin.Context, db *sql.DB, user models.User) {

	user.CreatedAt = time.Now()

	// Hash the password
	hashedPassword, err := lib.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.HashedPassword = hashedPassword

	query := `INSERT INTO users (username, email, password, age, created_at) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.QueryRow(query, user.Username, user.Email, user.HashedPassword, user.Age, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user_id": user.ID})

}

func Auth(c *gin.Context, db *sql.DB, user models.User) {

	// Hash the password
	hashedPassword, err := lib.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.HashedPassword = hashedPassword

	// Retrieve user from the database
	query := `SELECT id, password FROM users WHERE username = $1`
	err = db.QueryRow(query, user.Username).Scan(&user.ID, &user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password :455"})
		return
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password :555"})
		return
	}

	// Validate user and generate JWT token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username :": user.Username, "token : ": token})

}
