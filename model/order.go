package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	Status    string             `json:"status" bson:"status"`
	StatusInt int                `json:"statusInt" bson:"statusInt"`
	User      User               `json:"user" bson:"user"`
	Cart      map[string]int     `json:"cart" bson:"cart"`
}
