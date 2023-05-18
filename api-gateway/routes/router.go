package routes

import "github.com/gin-gonic/gin"

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use()
}
