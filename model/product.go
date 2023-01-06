package model

type Product struct {
	Code      string  `json:"code" bson:"code"`
	Name      string  `json:"name" bson:"name"`
	Origin    string  `json:"origin" bson:"origin"`
	Price     float64 `json:"price" bson:"price"`
	Limit     int     `json:"limit" bson:"limit"`
	CanOrder  bool    `json:"canOrder" bson:"canOrder"`
	CanView   bool    `json:"canView" bson:"canView"`
	CreatedAt int64   `json:"createdAt" bson:"createdAt"`
}
