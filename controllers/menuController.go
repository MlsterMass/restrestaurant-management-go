package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"restaurant/db"
	"restaurant/models"
	"time"
)

var menuCollection *mongo.Collection = db.OpenCollection(db.Client, "menu")

func GetMenus(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	result, err := menuCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the menus items"})
	}
	var allMenus []bson.M
	if err = result.All(c, &allMenus); err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, allMenus)
}

func GetMenu(ctx *gin.Context) {
	var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	menuId := ctx.Param("menu_id")
	var menu models.Menu

	err := foodCollection.FindOne(c, bson.M{"menu_id": menuId}).Decode(&menu)
	defer cancel()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food item"})
	}
	ctx.JSON(http.StatusOK, menu)
}
