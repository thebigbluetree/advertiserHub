package MongoDB

import (
	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)



func MongoConn() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	//err = client.Ping(context.TODO(), nil)
	//if err != nil{
	//	log.Fatal(err)
	//}
	return client
}
