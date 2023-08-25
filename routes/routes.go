package routes

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {
	authRoute(g)
	projectRoute(g)
	profileRoute(g)
	userRoute(g)
	adsRoute(g)
}
