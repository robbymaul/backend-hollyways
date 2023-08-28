package routes

import "github.com/gin-gonic/gin"

func RouteInit(r *gin.RouterGroup) {
	authRoute(r)
	branRoute(r)
	logoRoute(r)
	projectRoute(r)
	profileRoute(r)
	userRoute(r)
	adsRoute(r)
}
