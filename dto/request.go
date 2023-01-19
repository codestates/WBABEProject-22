package dto

type OrderCreate struct {
	User UserCreate `json:"user" bson:"user" binding:"required"`
	OrderUpdateCart
}

type OrderUpdateStatus struct {
	Status string `json:"status" bson:"status" binding:"required,oneof=Submitting Submitted Cooking Cooked Delivering Delivered Cancelled"`
}

type OrderUpdateCart struct {
	Cart map[string]int `json:"cart" bson:"cart" binding:"required" swaggertype:"object,integer" example:"productCode1:1,productCode2:1"`
}

type ProductCreate struct {
	Code string `json:"code" bson:"code" binding:"required" example:"bc01"`
	ProductUpdate
}

type ProductUpdate struct {
	Name     string  `json:"name" bson:"name" binding:"required" example:"Chicken burrito"`
	Origin   string  `json:"origin" bson:"origin" binding:"required" example:"Mexico"`
	Price    float64 `json:"price" bson:"price" binding:"required" example:"9.99"`
	Limit    int     `json:"limit" bson:"limit" binding:"required" example:"100"`
	CanOrder bool    `json:"canOrder" bson:"canOrder" binding:"required" example:"true"`
	CanView  bool    `json:"canView" bson:"canView" binding:"required" example:"true"`
}

type ReviewOrderCreate struct {
	Rating         float64               `json:"rating" bson:"rating" binding:"required,min=0,max=5" example:"4.5"`
	Comment        string                `json:"comment" bson:"comment" example:"Incredible!"`
	ReviewProducts []ReviewProductCreate `json:"reviewProducts" bson:"reviewProducts"`
}

type ReviewProductCreate struct {
	ProductCode string `json:"productCode" bson:"productCode" binding:"required" example:"bc01"`
	IsLiked     bool   `json:"isLiked" bson:"isLiked" binding:"required" example:"false"`
	Comment     string `json:"comment" bson:"comment" example:"Good!"`
}

type UserCreate struct {
	Username string        `json:"username" bson:"username" binding:"required,alphanum,max=30" example:"abc1"`
	Email    string        `json:"email" bson:"email" binding:"required,email" example:"abc1@gmail.com"`
	Phone    string        `json:"phone" bson:"phone" binding:"required,e164" example:"+821011112222"`
	Address  AddressCreate `json:"address" bson:"address" binding:"required"`
}

type AddressCreate struct {
	CountryCode        string `json:"countryCode" bson:"countryCode" binding:"required,iso3166_1_alpha3" example:"KOR"`
	AdministrativeArea string `json:"administrativeArea" bson:"administrativeArea" binding:"required" example:"Seoul"`
	Locality           string `json:"locality" bson:"locality" binding:"required" example:"Jongno-gu"`
	DependentLocality  string `json:"dependentLocality" bson:"dependentLocality" example:""`
	StreetAddress      string `json:"streetAddress" bson:"streetAddress" binding:"required" example:"Jong-ro 1"`
	PostalCode         string `json:"postalCode" bson:"postalCode" binding:"required,alphanum,len=5" example:"03154"`
}
