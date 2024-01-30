package auth

import (
	"api_web/web_profile/data"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Loginbody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json Loginbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser data.User
	data.Db.Where("username = ?", json.Username).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Username not Exist",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(json.Password))
	if err == nil {

		// expirationTime := time.Now().Add(1 * time.Hour).Unix()
		expirationTime := time.Now().Add(1 * time.Minute).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       existingUser.ID,
			"Fullname": existingUser.Fullname,
			"exp":      expirationTime,
		})

		hmacSampleSecret := []byte("login-myweb")
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			fmt.Println("Error signing the token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal Server Error"})
			return
		}

		// ตั้งค่า Header Authorization ใน Response
		c.Header("Authorization", tokenString)



		

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Login Successful",
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Login Failed"})
	}
}
