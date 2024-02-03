package auth

import (
	"api_web/web_profile/data"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type Registerbody struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Isuse    string `json:"isuse"` 
}



func Register (c *gin.Context){
	var json Registerbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating hash"})
		return
	}

	// เช็คว่ามี username อยู่แล้วหรือยัง
	var existingUser data.User
	if err := data.Db.Where("username = ?", json.Username).First(&existingUser).Error; err == nil {
		// Username already exists
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists"})
		return
	}

		// Create user in the database with isuse = 1
		user := data.User{
			Fullname: json.Fullname,
			Email:    json.Email,
			Username: json.Username,
			Password: string(hashPassword),
			Isuse:    "1",
		}

	result := data.Db.Create(&user)

	//ผลลัพธ์ error ต้องไม่มีและ 
	if result.Error == nil && result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":  "User created successfully",
			"status":   "success",
			// "register": user,
			"error":    nil,
		})
	} else {
		// Return error response StatusBadRequest = 400
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not created",
			"error":   result.Error,
		})
	}
}