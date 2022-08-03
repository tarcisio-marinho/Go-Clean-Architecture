package repository

import (
	"context"
	"go-clean-architecture/src/domain"
	"go-clean-architecture/src/infrastructure/repository/io"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IGetCustomerQuery interface {
	Get(clientId string) (domain.Customer, error)
}

type GetCustomerQuery struct {
	Database CreateDatabase
}

func (query GetCustomerQuery) Get(customerId string) (domain.Customer, error) {
	db, err := query.Database.Create()

	output := domain.Customer{}
	if err != nil {
		return domain.Customer{}, err
	}

	collection := db.Collection(query.Database.Collection)

	projection := io.Projection{
		Id:       1,
		Name:     1,
		BirthDay: 1,
	}

	filter := bson.M{
		"customer_id": customerId,
	}

	findOptions := options.FindOne().
		SetProjection(projection)

	entry := io.CustomerEntry{}
	err = collection.FindOne(context.TODO(), filter, findOptions).Decode(&entry)

	if err != nil {
		return output, err
	}

	return entry.ToOrderEvaluation(), nil
}
