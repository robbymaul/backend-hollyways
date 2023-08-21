package routes

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {
	AuthRoute(g)
	ProjectRoute(g)
	UserRoute(g)
}
