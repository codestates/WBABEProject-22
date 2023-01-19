package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"oos/db"
	"oos/dto"
	"oos/model"
)

func CreateProduct(ctx context.Context, params dto.ProductCreate) (*mongo.InsertOneResult, error) {
	product := model.Product{
		CreatedAt:  time.Now().UnixMicro(),
		UpdatedAt:  time.Now().UnixMicro(),
		UserOrders: map[string]int{},
		ProductView: model.ProductView{
			ProductCreate: dto.ProductCreate{
				Code: params.Code,
				ProductUpdate: dto.ProductUpdate{
					Name:     params.Name,
					Origin:   params.Origin,
					Price:    params.Price,
					Limit:    params.Limit,
					CanOrder: params.CanOrder,
					CanView:  true,
				},
			},
		},
	}

	result, err := db.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ListProducts(ctx context.Context, sortBy string) ([]model.ProductView, error) {
	setStage := bson.D{{Key: "$set", Value: bson.M{
		"ratings":  bson.M{"$divide": bson.A{"$productquery.ratingSum", "$productquery.reviewCount"}},
		"reorders": bson.M{"$sum": "$userOrders"},
		"likes":    "$productquery.likeCount",
		"time":     "$createdAt",
	}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{sortBy: -1}}}

	cursor, err := db.ProductCollection.Aggregate(ctx, mongo.Pipeline{setStage, sortStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.ProductView
	for cursor.Next(ctx) {
		var product model.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}
		productView := model.ProductView{
			RatingSum:   product.RatingSum,
			LikeCount:   product.LikeCount,
			ReviewCount: product.ReviewCount,
			ProductCreate: dto.ProductCreate{
				Code: product.Code,
				ProductUpdate: dto.ProductUpdate{
					Name:     product.Name,
					Origin:   product.Origin,
					Price:    product.Price,
					Limit:    product.Limit,
					CanOrder: product.CanOrder,
					CanView:  product.CanView,
				},
			},
		}
		products = append(products, productView)
	}

	return products, nil
}

func GetProduct(ctx context.Context, productCode string) (*model.Product, error) {
	filter := bson.M{"productquery.productcreate.code": productCode}

	var product model.Product
	if err := db.ProductCollection.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProduct(ctx context.Context, productCode string, product dto.ProductUpdate) (*mongo.UpdateResult, error) {
	filter := bson.M{"productquery.productcreate.code": productCode}
	update := bson.M{"$set": bson.M{
		"productquery.productcreate.productupdate.name":     product.Name,
		"productquery.productcreate.productupdate.origin":   product.Origin,
		"productquery.productcreate.productupdate.price":    product.Price,
		"productquery.productcreate.productupdate.limit":    product.Limit,
		"productquery.productcreate.productupdate.canOrder": product.CanOrder,
		"productquery.productcreate.productupdate.canView":  product.CanView,
		"updatedAt": time.Now().UnixMicro(),
	}}

	result, err := db.ProductCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount != 1 {
		err = errors.New("no match to update")
		return nil, err
	}

	return result, nil
}

func DeleteProduct(ctx context.Context, productCode string) (*mongo.UpdateResult, error) {
	filter := bson.M{"productquery.productcreate.code": productCode}
	update := bson.M{"$set": bson.M{
		"productquery.productcreate.productupdate.canView": false,
		"updatedAt": time.Now().UnixMicro(),
	}}

	result, err := db.ProductCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount != 1 {
		err = errors.New("no match to delete")
		return nil, err
	}

	return result, nil
}
