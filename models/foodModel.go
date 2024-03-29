package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Price     *float64           `json:"price" validate:"required"`
	FoodImage *string            `json:"food_image" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" `
	UpdatedAt time.Time          `json:"updatedAt"`
	FoodId    string             `json:"foodId"`
	MenuId    *string            `json:"menuId" validate:"required"`
}
