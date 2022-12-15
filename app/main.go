package main

import (
	"go-todo-app/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"msg": "pong!"}) })

	r.Run(":7790")
}
