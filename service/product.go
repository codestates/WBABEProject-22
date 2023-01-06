package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/db"
	"oos/dto"
	"oos/model"
)

func CreateProduct(ctx context.Context, params dto.ProductCreate) (*mongo.InsertOneResult, error) {
	product := model.Product{
		Code:      params.Code,
		Name:      params.Name,
		Origin:    params.Origin,
		Price:     params.Price,
		Limit:     params.Limit,
		CanOrder:  params.CanOrder,
		CanView:   true,
		CreatedAt: time.Now().UnixMicro(),
	}

	result, err := db.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProducts(ctx context.Context) ([]model.Product, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"code": 1})

	cursor, err := db.ProductCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func GetProduct(ctx context.Context, productCode string) (*model.Product, error) {
	filter := bson.M{"code": productCode}

	var product model.Product
	err := db.ProductCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProduct(ctx context.Context, productCode string, product dto.ProductUpdate) (*mongo.UpdateResult, error) {
	filter := bson.M{"code": productCode}
	update := bson.M{"$set": bson.M{
		"name":     product.Name,
		"origin":   product.Origin,
		"price":    product.Price,
		"limit":    product.Limit,
		"canOrder": product.CanOrder,
		"canView":  product.CanView,
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
	filter := bson.M{"code": productCode}
	update := bson.M{"$set": bson.M{"canView": false}}

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
