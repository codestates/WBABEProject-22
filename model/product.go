package model

import "oos/dto"

type Product struct {
	CreatedAt  int64          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  int64          `json:"updatedAt" bson:"updatedAt"`
	UserOrders map[string]int `json:"userOrders" bson:"userOrders"`
	ProductView
}

type ProductView struct {
	RatingSum   float32 `json:"ratingSum" bson:"ratingSum"`
	LikeCount   int     `json:"likeCount" bson:"likeCount"`
	ReviewCount int     `json:"reviewCount" bson:"reviewCount"`
	dto.ProductCreate
}
