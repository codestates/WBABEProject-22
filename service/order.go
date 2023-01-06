package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/db"
	"oos/dto"
	"oos/model"
)

func CreateOrder(ctx context.Context, params dto.OrderCreate) (*mongo.InsertOneResult, error) {
	order := model.Order{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().UnixMicro(),
		Status:    "Submitting",
		StatusInt: model.OrderStatus["Submitting"],
		User:      params.User,
		Cart:      params.Cart,
	}

	result, err := db.OrderCollection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

func GetOrders(ctx context.Context) ([]model.Order, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.M{"createdAt": 1})

	cursor, err := db.OrderCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []model.Order
	for cursor.Next(ctx) {
		var order model.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrdersActive(ctx context.Context, username string) ([]model.Order, error) {
	filter := bson.M{"$and":
		bson.A{
			bson.M{"user.username": username},
			bson.M{"statusInt": bson.M{"$lt": 5}},
		},
	}
	opts := options.Find().SetSort(bson.M{"createdAt": 1})

	cursor, err := db.OrderCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []model.Order
	for cursor.Next(ctx) {
		var order model.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrdersHistory(ctx context.Context, username string) ([]model.Order, error) {
	filter := bson.M{"$and":
		bson.A{
			bson.M{"user.username": username},
			bson.M{"statusInt": bson.M{"$gte": 5}},
		},
	}
	opts := options.Find().SetSort(bson.M{"createdAt": 1})

	cursor, err := db.OrderCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []model.Order
	for cursor.Next(ctx) {
		var order model.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrder(ctx context.Context, orderId string) (*model.Order, error) {
	orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": orderIdObject}

	var order model.Order
	err := db.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrderStatus(ctx context.Context, orderId string) (*string, error) {
	orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": orderIdObject}

	var order model.Order
	err := db.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order.Status, nil
}

func UpdateOrderStatus(ctx context.Context, orderId string, params dto.OrderUpdateStatus) (*mongo.UpdateResult, error) {
	orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": orderIdObject}
	update := bson.M{"$set": bson.M{"status": params.Status, "statusInt": model.OrderStatus[params.Status]}}

	result, err := db.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount != 1 {
		return nil, errors.New("no match to update")
	}
	
	return result, nil
}

func UpdateOrderItems(ctx context.Context, orderId string, params dto.OrderUpdateCart) (*mongo.UpdateResult, error) {
	orderStatus, err := GetOrderStatus(ctx, orderId)
	if err != nil {
		return nil, err
	}
	if model.OrderStatus[*orderStatus] >= model.OrderStatus["Delivering"] {
		return nil, errors.New("order change not allowed at this stage")
	}

	orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": orderIdObject}

	var order model.Order
	err = db.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	for productCode, quantity := range params.Cart {
		order.Cart[productCode] = quantity
	}

	update := bson.M{"$set": bson.M{"cart": order.Cart}}

	result, err := db.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount != 1 {
		return nil, errors.New("no match to update")
	}

	return result, nil
}

func DeleteOrderItems(ctx context.Context, orderId string, params []string) (*mongo.UpdateResult, error) {
	orderStatus, err := GetOrderStatus(ctx, orderId)
	if err != nil {
		return nil, err
	}
	if model.OrderStatus[*orderStatus] >= model.OrderStatus["Cooking"] {
		return nil, errors.New("order change not allowed at this stage")
	}

	orderIdObject, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": orderIdObject}

	var order model.Order
	err = db.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return nil, err
	}

	for _, productCode := range params {
		delete(order.Cart, productCode)
	}

	update := bson.M{"$set": bson.M{"cart": order.Cart}}

	result, err := db.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount != 1 {
		return nil, errors.New("no match to delete")
	}

	return result, nil
}