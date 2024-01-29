package main

import (
	auth_controller "api_web/web_profile/controllers/auth"
	user_controller "api_web/web_profile/controllers/user"
	"api_web/web_profile/data"

	// _"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Username string
	Password string
	// Role     string
}

func main() {
	data.Userdb()
	r := gin.New()
	r.Use(cors.Default())
	r.GET("/users", user_controller.Readall)

	r.POST("/register", auth_controller.Register)
	r.POST("/login", auth_controller.Login)

	r.Run("localhost:8880")
}
