package repository

import (
	"context"
	"go-clean-architecture/src/api/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type CreateDatabase struct {
	DbConfig   configs.DatabaseConfig
	Database   string
	Collection string
}

func (db CreateDatabase) Create() (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DbConfig.GetURILocal()))

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(db.Database), nil
}
