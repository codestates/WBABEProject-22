package model

import "oos/dto"

type ReviewOrder struct {
	OrderID  string `json:"orderID" bson:"orderID"`
	Username string `json:"username" bson:"username"`
	dto.ReviewOrderCreate
}

type ReviewProduct struct {
	dto.ReviewProductCreate
}

type ReviewView struct {
	Username string  `json:"username" bson:"username"`
	Rating   float64 `json:"rating" bson:"rating"`
	dto.ReviewProductCreate
}

// References
// https://www.mongodb.com/docs/manual/tutorial/model-embedded-one-to-one-relationships-between-documents/
// https://www.mongodb.com/developer/products/mongodb/mongodb-schema-design-best-practices/
