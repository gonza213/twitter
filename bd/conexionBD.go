package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN
var MongoCN = ConectarBD()
var clientsOptions = options.Client().ApplyURI("mongodb+srv://root:gonza1991@cluster0-gqowt.mongodb.net/test?retryWrites=true&w=majority")

//ConectarBD
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientsOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la BD")
	return client
}

// ChequeoConnection
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
