package restaurant_management

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"restaurant/db"
	"restaurant/middleware"
	"restaurant/routes"
)

var foodCollection *mongo.Collection = db.OpenCollection(db.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
