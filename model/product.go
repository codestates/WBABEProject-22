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
	/* [코드리뷰]
	 * DB와 관련된 데이터 모델들을 보통 CreateAt 과 ModifyAt과 같은 update 시기에 따른 데이터도 가져가고 있습니다.
	 * ModifyAt 속성을 추가하여, 데이터 정합성을 고려한 포인트를 만들어두시는 코드 방식도 추천드립니다.
	 */
}
