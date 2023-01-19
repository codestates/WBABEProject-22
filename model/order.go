package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oos/dto"
)

var OrderStatus = map[string]int{
	"Submitting": 0,
	"Submitted":  1,
	"Cooking":    2,
	"Cooked":     3,
	"Delivering": 4,
	"Delivered":  5,
	"Cancelled":  6,
}

type Order struct {
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Status    string             `json:"status" bson:"status" binding:"required,oneof=Submitting Submitted Cooking Cooked Delivering Delivered Cancelled"`
	User      dto.UserCreate     `json:"user" bson:"user"`
	Cart      map[string]int     `json:"cart" bson:"cart" binding:"required" swaggertype:"object,integer" example:"productCode1:1,productCode2:1"`
}
