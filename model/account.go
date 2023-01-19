package model

type Token struct {
	UserRole string `json:"userRole" bson:"userRole"`
	JwtToken string `json:"jwtToken" bson:"jwtToken"`
}
