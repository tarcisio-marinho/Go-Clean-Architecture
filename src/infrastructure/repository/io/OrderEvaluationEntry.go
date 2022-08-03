package io

import "go-clean-architecture/src/domain"

type CustomerEntry struct {
	Id       string `bson:"customer_id"`
	Name     string `bson:"customer_name"`
	BirthDay string `bson:"customer_birthday"`
}

func (entry CustomerEntry) ToOrderEvaluation() domain.Customer {
	return domain.Customer{
		Id:        entry.Id,
		Name:      entry.Name,
		BirthDate: entry.BirthDay,
	}
}
