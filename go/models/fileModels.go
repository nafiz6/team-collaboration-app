package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskFile struct {
	ID           primitive.ObjectID
	FileName     string
	Url          string
	TaskId       primitive.ObjectID
	Date_created primitive.DateTime
}

type WorkspaceFile struct {
	ID           primitive.ObjectID
	FileName     string
	Url          string
	WorkspaceId  primitive.ObjectID
	Date_created primitive.DateTime
}
