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
/* [코드리뷰]
 * 주문상태에 대한 Status를 key-value 형식으로 잘 관리해주고 있는 훌륭한 코드입니다.
 * 이런 코딩 스타일을 통해, 앞으로 주문 상태가 복잡해지는 상황에서도 충분한 대응이 가능하게 됩니다.
 * 잘 짜주셨습니다.
 */

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	Status    string             `json:"status" bson:"status"`
	StatusInt int                `json:"statusInt" bson:"statusInt"`
	User      User               `json:"user" bson:"user"`
	Cart      map[string]int     `json:"cart" bson:"cart"`
}
