package io

type Projection struct {
	Id       int `bson:"customer_id"`
	Name     int `bson:"customer_name"`
	BirthDay int `bson:"customer_birthday"`
}
