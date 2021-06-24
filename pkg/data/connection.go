package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tonoy30/echo-go/pkg/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewConnection(settings *settings.Settings) Connection {
	uri := fmt.Sprintf("%s://%s:%s/%s", settings.DBDriver, settings.DBHost, settings.DBPort, settings.DBName)

	credentials := options.Credential{
		Username: settings.DBUser,
		Password: settings.DBPassword,
	}

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to mongodb using uri: %s", uri)

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

func (c Connection) Disconnect() {
	log.Println("Database disconnecting...")
	_ = c.Client.Disconnect(c.ctx)
}
