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
	} else {
		// If the token is valid, proceed with processing
		var users []data.User
		data.Db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "This is all user", "data-users": users})
	}
}

func GetPerson(c *gin.Context) {
	Username := c.Params.ByName("username")
	var person data.User
	if err := data.Db.Where("username = ?", Username).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func UpdatePerson(c *gin.Context) {
	var person data.User
	Username := c.Params.ByName("username")
	if err := data.Db.Where("username = ?", Username).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	data.Db.Save(&person)
	c.JSON(200, person)
}



func DeletePerson(c *gin.Context) {
	Username := c.Params.ByName("username") //ลบตาม username
	var person data.User 

	// ค้นหา username ที่ต้องการลบ
	if err := data.Db.Where("username = ?", Username).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// ทำการอัพเดทค่า isuse เป็น 0
	result := data.Db.Model(&person).Update("isuse", 0)

	if result.Error == nil && result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "status": "success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not deleted", "error": result.Error})
	}
}
