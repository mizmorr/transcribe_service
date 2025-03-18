package delivery

import "github.com/gin-gonic/gin"

type controller interface {
	Transcribe(g *gin.Context)
}

func NewRouter(router *gin.Engine, c controller) {
	router.Use(gin.Recovery())

	router.Use(gin.Logger())

	publicRoutes := router.Group("/api/v1")
	{
		publicRoutes.POST("/transcribe", c.Transcribe)
	}
}
