package model

type ReviewOrder struct {
	OrderID        string          `json:"orderId" bson:"orderId"`
	Username       string          `json:"username" bson:"username"`
	Rating         float64         `json:"rating" bson:"rating"`
	Comment        string          `json:"comment" bson:"comment"`
	ReviewProducts []ReviewProduct `json:"productReviews" bson:"productReviews"`
}

type ReviewProduct struct {
	ProductCode string `json:"productCode" bson:"productCode" binding:"required" example:"bc01"`
	Recommend   bool   `json:"recommend" bson:"recommend" binding:"required" example:"false"`
	Comment     string `json:"comment" bson:"comment" example:"Good!"`
}

type ReviewProductQuery struct {
	Username    string  `json:"username" bson:"username" binding:"required,alphanum,max=30" example:"abc1"`
	Rating      float64 `json:"rating" bson:"rating" binding:"required,min=0,max=5" example:"4.5"`
	ProductCode string  `json:"productCode" bson:"productCode" binding:"required" example:"bc01"`
	Recommend   bool    `json:"recommend" bson:"recommend" binding:"required" example:"false"`
	Comment     string  `json:"comment" bson:"comment" example:"Good!"`
}

// References
// https://www.mongodb.com/docs/manual/tutorial/model-embedded-one-to-one-relationships-between-documents/
// https://www.mongodb.com/developer/products/mongodb/mongodb-schema-design-best-practices/
