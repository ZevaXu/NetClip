package router

import (
	"NetClip2/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.SetTrustedProxies([]string{"100.125.192.53:8080"})
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/login", controller.Login)
	e.POST("/login", controller.AddUser)

	e.GET("/:username/:password", controller.UserIndex)
	e.POST("/s", controller.UpdateUserContent)
	e.POST("/d", controller.DeleteUser)

	e.Run("0.0.0.0:8080")
}
