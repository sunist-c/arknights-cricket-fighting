package router

import "github.com/gin-gonic/gin"

var engine = gin.Default()

func GetEngine() *gin.Engine {
	return engine
}
