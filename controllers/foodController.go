package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"restaurant/db"
	"restaurant/models"
	"time"
)

var foodCollection *mongo.Collection = db.OpenCollection(db.Client, "food")
var validate = validator.New()

func GetFoods(ctx *gin.Context) {

}

func GetFood(ctx *gin.Context) {
	var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	foodId := ctx.Param("food_id")
	var food models.Food

	err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food)
	defer cancel()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food item"})
	}
	ctx.JSON(http.StatusOK, food)
}

func CreateFood(ctx *gin.Context) {
	var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var menu models.Menu
	var food models.Food

	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validateErr := validate.Struct(food)
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
	}
	err := menuCollection.FindOne(c, bson.M{"menu_id": food.MenuId}).Decode(&menu)
	defer cancel()
	if err != nil {
		msg := fmt.Sprintf("menu don't find")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	food.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.ID = primitive.NewObjectID()
	food.FoodId = food.ID.Hex()
	var num = toFixed(*food.Price, 2)
	food.Price = &num

	result, insertErr := foodCollection.InsertOne(c, food)
	if insertErr != nil {
		msg := fmt.Sprintf("Food item was not created")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	ctx.JSON(http.StatusOK, result)
}

func round(num float64) int {

}

func toFixed(num float64, precission int) float64 {

}
func UpdateFood(ctx *gin.Context) {

}
