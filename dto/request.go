package dto

import "oos/model"

type ProductCreate struct {
	Code     string  `json:"code" bson:"code" binding:"required" example:"bc01"`
	Name     string  `json:"name" bson:"name" binding:"required" example:"Chicken burrito"`
	Origin   string  `json:"origin" bson:"origin" binding:"required" example:"Mexico"`
	Price    float64 `json:"price" bson:"price" binding:"required" example:"9.99"`
	Limit    int     `json:"limit" bson:"limit" binding:"required" example:"100"`
	CanOrder bool    `json:"canOrder" bson:"canOrder" binding:"required" example:"true"`
}

type ProductUpdate struct {
	Name     string  `json:"name" bson:"name" binding:"required"`
	Origin   string  `json:"origin" bson:"origin" binding:"required"`
	Price    float64 `json:"price" bson:"price" binding:"required"`
	Limit    int     `json:"limit,omitempty" bson:"limit" binding:"required"`
	CanOrder bool    `json:"canOrder" bson:"canOrder" binding:"required"`
	CanView  bool    `json:"canView" bson:"canView" binding:"required"`
}

type OrderCreate struct {
	User model.User     `json:"user" bson:"user" binding:"required"`
	Cart map[string]int `json:"cart" bson:"cart" binding:"required"`
}

type OrderUpdateStatus struct {
	Status string `json:"status" bson:"status" binding:"required"`
}

type OrderUpdateCart struct {
	Cart map[string]int `json:"cart" bson:"cart" binding:"required"`
}

type ReviewOrderCreate struct {
	Rating         float64               `json:"rating" bson:"rating" binding:"required,min=0,max=5" example:"4.5"`
	Comment        string                `json:"comment" bson:"comment" example:"Incredible!"`
	ReviewProducts []model.ReviewProduct `json:"productReviews" bson:"productReviews"`
}
