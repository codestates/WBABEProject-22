package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/db"
	"oos/dto"
	"oos/model"
)

func CreateReview(ctx context.Context, orderId string, params dto.ReviewOrderCreate) (*mongo.InsertOneResult, error) {
	// orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	// filter := bson.M{"_id": orderIdObject}

	// var order model.Order
	// err := db.OrderCollection.FindOne(ctx, filter).Decode(&order)
	// if err != nil {
	// 	return nil, err
	// }

	order, err := GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	newReview := model.ReviewOrder{
		OrderID:        orderId,
		Username:       order.User.Username,
		Rating:         params.Rating,
		Comment:        params.Comment,
		ReviewProducts: params.ReviewProducts,
	}

	result, err := db.ReviewCollection.InsertOne(ctx, newReview)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetReviews(ctx context.Context) ([]model.ReviewOrder, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"orderId": 1})

	cursor, err := db.ReviewCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []model.ReviewOrder
	for cursor.Next(ctx) {
		var review model.ReviewOrder
		err = cursor.Decode(&review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func GetReview(ctx context.Context, productCode string) ([]model.ReviewProductQuery, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"orderId": 1})

	cursor, err := db.ReviewCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []model.ReviewProductQuery
	for cursor.Next(ctx) {
		var orderReview model.ReviewOrder
		err = cursor.Decode(&orderReview)
		if err != nil {
			return nil, err
		}
		for _, productReview := range orderReview.ReviewProducts {
			if productReview.ProductCode == productCode {
				newProductReview := model.ReviewProductQuery{
					Username:    orderReview.Username,
					Rating:      orderReview.Rating,
					ProductCode: productReview.ProductCode,
					Recommend:   productReview.Recommend,
					Comment:     productReview.Comment,
				}
				reviews = append(reviews, newProductReview)
			}
		}
	}

	return reviews, nil
}
