package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods)
	incomingRoutes.GET("/foods/:foods_id", controllers.GetFood)
	incomingRoutes.POST("/foods", controllers.CreateFood)
	incomingRoutes.POST("/foods/:food_id", controllers.UpdateFood)
}
