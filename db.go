package EComDB

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnector - The struct that defines the database connector
type DBConnector struct {
	Client *mongo.Client
}

// NewDBConnector - Create a new instance of a database conenctor
func NewDBConnector() *DBConnector {
	return &DBConnector{}
}

// Init - Initialize the database connector
func (dbc *DBConnector) Init() bool {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	dbc.Client = client
	return true
}
