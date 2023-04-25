package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nafees-dev/packaform_test/controller"
)

type OrderRouteController struct {
	orderController controller.OrderController
}

func NewRouteOrder(orderController controller.OrderController) OrderRouteController {
	return OrderRouteController{orderController}
}

func (order *OrderRouteController) OrderRoute(rg *gin.RouterGroup) {
	router := rg.Group("orders")
	// router.Use(middleware.DeserializeUser())
	router.GET("/", order.orderController.GetAllOrders)
	// router.GET("/:postId", pc.postController.FindPostById)
}
