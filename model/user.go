package model

type User struct {
	Username string  `json:"username" bson:"username" binding:"required,alphanum,max=30" example:"abc1"`
	Email    string  `json:"email" bson:"email" binding:"required,email" example:"abc1@gmail.com"`
	Phone    string  `json:"phone" bson:"phone" binding:"required,e164" example:"+821011112222"`
	Address  Address `json:"address" bson:"address" binding:"required"`
}

type Address struct {
	CountryCode        string `json:"countryCode" bson:"countryCode" binding:"required,iso3166_1_alpha3" example:"KOR"`
	AdministrativeArea string `json:"administrativeArea" bson:"administrativeArea" binding:"required" example:"Seoul"`
	Locality           string `json:"locality" bson:"locality" binding:"required" example:"Jongno-gu"`
	DependentLocality  string `json:"dependentLocality" bson:"dependentLocality"`
	StreetAddress      string `json:"streetAddress" bson:"streetAddress" binding:"required" example:"Jong-ro 1"`
	PostalCode         string `json:"postalCode" bson:"postalCode" binding:"required,alphanum,len=5" example:"03154"`
}

// References
// https://github.com/nyaruka/phonenumbers
// https://phonenumbers.temba.io/
// https://github.com/Boostport/address
// https://chromium-i18n.appspot.com/ssl-address