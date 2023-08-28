package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func transactionRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerTransaction(repo)

	r.POST("/transaction", middleware.Auth(), handler.CreateTransaction)
	r.GET("/transactions", middleware.Auth(), handler.FindTransaction)
	r.GET("/transaction/:id", middleware.Auth(), handler.GetTransaction)
	r.GET("/payment/:id", handler.GetPayment)
	r.POST("/notification", handler.Notification)
}
