package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserLogin struct {
	Username string
	Password string
}

type UserRegistration struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string
	Password    string
	Username    string
	Date_joined primitive.DateTime
	Dp          string
	Bio         string
}
