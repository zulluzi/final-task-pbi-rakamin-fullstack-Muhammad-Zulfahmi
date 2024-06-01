package router

import (
	"log"
	"net/http"

	"BTPNS_Fulldev/controllers"
	"BTPNS_Fulldev/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        log.Println("GET / called")
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to BTPNS API",
        })
    })

    r.POST("/users/register", func(c *gin.Context) {
        log.Println("POST /users/register called")
        controllers.Register(c)
    })
    r.POST("/users/login", func(c *gin.Context) {
        log.Println("POST /users/login called")
        controllers.Login(c)
    })

    auth := r.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.POST("/photos", func(c *gin.Context) {
            log.Println("POST /photos called")
            controllers.AddPhoto(c)
        })
        auth.DELETE("/photos/:photoId", func(c *gin.Context) {
            log.Println("DELETE /photos/:photoId called")
            controllers.DeletePhoto(c)
        })
    }

    log.Println("Router setup complete")
    return r
}
