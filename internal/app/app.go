package app

import "github.com/gin-gonic/gin"

type Router struct {
	router *gin.Engine
}

func InitRouter() Router {
	r := gin.Default()
	return Rout
}
