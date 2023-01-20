package model

type ReviewOrder struct {
	OrderID        string          `json:"orderId" bson:"orderId"`
	Username       string          `json:"username" bson:"username"`
	Rating         float64         `json:"rating" bson:"rating"`
	Comment        string          `json:"comment" bson:"comment"`
	ReviewProducts []ReviewProduct `json:"productReviews" bson:"productReviews"`
	/* [코드리뷰]
	 * 객체지향 프로그래밍 성격의 대표적인 모습을 담고 있는 좋은 코드입니다.
	 * ReviewProducts를 관리하며, ReviewProducts는 여러개의 ReviewProduct가 담길 수 있는 코드입니다.
	 * 이러한 코딩 방식을 통해 우선 코드에 대한 가독성이 명확하고, 
	 * 추후 유지보수 하기에도 어려움이 없는 코드라고 생각됩니다. 잘 만들어주셨습니다.
	 * 
	 * 추가로 rating에 있어서 min, max 값을 지정해두는 방법은 어떠실까요?
	 * 해당 점수를 지표화하여, 추후 데이터 분석을 위한 자료로 만들려면
	 * custom한 숫자형 데이터인 경우 range를 지정해주는 방법이 필요할 것으로 보여집니다.
	 */
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
