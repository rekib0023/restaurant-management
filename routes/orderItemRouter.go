package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ordersItems", controller.GetOrderItems())
	incomingRoutes.GET("/ordersItems/:orderItemId", controller.GetOrderItem())
	incomingRoutes.GET("/ordersItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/ordersItems/:orderItemId", controller.UpdateOrderItem())
}
