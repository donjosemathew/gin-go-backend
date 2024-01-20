package main

import (
	"gilab.com/progmaticreviews/golang-gin-poc/controller"
	"gilab.com/progmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (videoService service.VideoService=service.New()
	videController controller.VideoController=controller.New(videoService)
)
func main() {
	server := gin.Default()
	server.GET("/test",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"Hello",
		})
	})
	server.GET("/videos",func(ctx *gin.Context){
		ctx.JSON(200,videController.FindAll())
	})

	server.POST("/videos",func(ctx *gin.Context){
		ctx.JSON(200,videController.Save(ctx))
	})
	server.Run(":8080")
}