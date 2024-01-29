package user

import (
	"api_web/web_profile/data"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Readall(c *gin.Context) {
	hmacSampleSecret := []byte("login-myweb")
	header := c.Request.Header.Get("Authorization")
	fullToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))

	if fullToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "คุณยังไม่มี Token กรุณา Login !!!"})
		return
	}

	token, err := jwt.Parse(fullToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Provide the secret key for validation
		return hmacSampleSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Token หมดอายุ !!!"})
		return
	}else{
		// If the token is valid, proceed with processing
		var users []data.User
		data.Db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "This is all user", "data-users": users})
		// c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "This is all user", "data-users": users, "header": header, "token": fullToken, "Full-token": fullToken})
	}

	
}
