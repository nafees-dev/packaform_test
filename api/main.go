package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nafees-dev/packaform_test/controller"
	"github.com/nafees-dev/packaform_test/initializers"
	"github.com/nafees-dev/packaform_test/routes"
	// "github.com/nafees-dev/packaform_test/controllers"
	// "github.com/nafees-dev/packaform_test/gin-gonic/gin"
	// "github.com/nafees-dev/packaform_test/initializers"
	// "github.com/nafees-dev/packaform_test/routes"
)

var (
	server *gin.Engine // AuthController      controllers.AuthController
	// AuthRouteController routes.AuthRouteController
	// UserController      controllers.UserController
	// UserRouteController routes.UserRouteController
	// PostController      controllers.PostController
	OrderController   controller.OrderController
	OrderRouteContext routes.OrderRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	// AuthController = controllers.NewAuthController(initializers.DB)
	// AuthRouteController = routes.NewAuthRouteController(AuthController)

	// UserController = controllers.NewUserController(initializers.DB)
	// UserRouteController = routes.NewRouteUserController(UserController)

	OrderController = controller.NewOrderController(initializers.DB)
	OrderRouteContext = routes.NewRouteOrder(OrderController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// AuthRouteController.AuthRoute(router)
	// UserRouteController.UserRoute(router)
	OrderRouteContext.OrderRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
