package EComDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBConnector struct {
	Client *mongo.Client
}

func (dbc *DBConnector) Init(dbtype string) bool {
	if dbtype == "mongodb" {
		clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		dbc.Client = client
		return true
	}
	return false
}
