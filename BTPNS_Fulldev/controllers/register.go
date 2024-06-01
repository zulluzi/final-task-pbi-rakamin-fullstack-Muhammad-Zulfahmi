package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    log.Println("Register controller called")
    c.JSON(http.StatusOK, gin.H{
        "message": "User registered successfully",
    })
}

func Login(c *gin.Context) {
    log.Println("Login controller called")
    c.JSON(http.StatusOK, gin.H{
        "message": "User logged in successfully",
    })
}

func AddPhoto(c *gin.Context) {
    log.Println("AddPhoto controller called")
    c.JSON(http.StatusOK, gin.H{
        "message": "Photo added successfully",
    })
}

func DeletePhoto(c *gin.Context) {
    log.Println("DeletePhoto controller called")
    c.JSON(http.StatusOK, gin.H{
        "message": "Photo deleted successfully",
    })
}
