package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"oos/config"
)

var DB *mongo.Client

var ProductCollection *mongo.Collection
var OrderCollection *mongo.Collection
var ReviewCollection *mongo.Collection

func ConnectDB(cfg *config.Config) {
	cf := cfg.DB
	uri := cf["host"]
	databaseName := cf["name"]
	
	DB = getDatabase(uri)
	
	ProductCollection = GetCollection(DB, databaseName, "products")
	OrderCollection = GetCollection(DB, databaseName, "orders")
	ReviewCollection = GetCollection(DB, databaseName, "reviews")
	
	// Product codes should be unique.
	_, err := ProductCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic(err)
		/* [코드리뷰]
		 * panic 함수를 사용하여, 각 상황에 대한 에러를 메세지와 함께 잘 정리해주셨습니다.
		 * 그러나 panic 함수 이외의 구문은 더 이상 코드가 동작하지 않게 되는데요,
		 * 실제로 운영되는 API 서비스라면, 서비스가 특정 상황으로 인해 종료되는 일은 발생하지 말아야 합니다.
		 * DB가 접근되지 않는 상황이라면 Critical한 상황이겠지만, 
		 * 예외처리를 통해 적절한 메세지를 개발자에게 전달하는 방식으로 수정해보시는 방법은 어떠실까요?
		 */
	}
}

func getDatabase(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}

func GetCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
