package routes

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {
	AuthRoute(g)
	testRoute(g)
	UserRoute(g)
}

func testRoute(g *gin.RouterGroup) {
	g.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
	})
}
