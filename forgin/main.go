package main

import (
	"one/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")
	router.Static("/static", "./static")

	router.GET("/", handler.HomePage)
	router.GET("/signup", handler.SignupPage)
	router.POST("/signuppost", handler.SignupPost)
	router.GET("/login", handler.LoginPage)
	router.POST("/loginpost", handler.Postmethod)
	router.GET("/logout", handler.Logout)
	router.Run(":8080")

}
