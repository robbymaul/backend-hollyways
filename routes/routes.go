package routes

import "github.com/gin-gonic/gin"

func RouteInit(r *gin.RouterGroup) {
	adsRoute(r)
	authRoute(r)
	branRoute(r)
	logoRoute(r)
	projectRoute(r)
	profileRoute(r)
	transactionRoute(r)
	userRoute(r)
}
