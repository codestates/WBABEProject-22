package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/db"
	"oos/dto"
	"oos/model"
)

func CreateReview(ctx context.Context, orderID string, params dto.ReviewOrderCreate) (*mongo.InsertOneResult, error) {
	order, err := GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}

	for _, reviewProduct := range params.ReviewProducts {
		like := 0
		if reviewProduct.IsLiked {
			like = 1
		}
		filter := bson.M{"productquery.productcreate.code": reviewProduct.ProductCode}
		update := bson.M{"$inc": bson.M{
			"userOrders." + order.User.Username: 1,
			"productquery.reviewCount":          1,
			"productquery.ratingSum":            params.Rating,
			"productquery.likeCount":            like,
		}}
		
		result, err := db.ProductCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
		if result.MatchedCount != 1 {
			return nil, errors.New("no match to update")
		}
	}

	review := model.ReviewOrder{
		OrderID:  orderID,
		Username: order.User.Username,
		ReviewOrderCreate: dto.ReviewOrderCreate{
			Rating:         params.Rating,
			Comment:        params.Comment,
			ReviewProducts: params.ReviewProducts,
		},
	}

	result, err := db.ReviewCollection.InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ListReviews(ctx context.Context) ([]model.ReviewOrder, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"orderID": 1})

	cursor, err := db.ReviewCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []model.ReviewOrder
	for cursor.Next(ctx) {
		var review model.ReviewOrder
		if err = cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func ListReviewsProduct(ctx context.Context, productCode string) ([]model.ReviewView, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"orderID": 1})

	cursor, err := db.ReviewCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []model.ReviewView
	for cursor.Next(ctx) {
		var reviewOrder model.ReviewOrder
		if err = cursor.Decode(&reviewOrder); err != nil {
			return nil, err
		}
		for _, reviewProduct := range reviewOrder.ReviewProducts {
			if reviewProduct.ProductCode == productCode {
				newReviewProduct := model.ReviewView{
					Username: reviewOrder.Username,
					Rating:   reviewOrder.Rating,
					ReviewProductCreate: dto.ReviewProductCreate{
						ProductCode: reviewProduct.ProductCode,
						IsLiked:     reviewProduct.IsLiked,
						Comment:     reviewProduct.Comment,
					},
				}
				reviews = append(reviews, newReviewProduct)
			}
		}
	}

	return reviews, nil
}
