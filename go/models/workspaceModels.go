package models

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
    ID primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
    Name string
    Workspaces []Workspace
}
type Workspace struct {
	Name string
	ID primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Users []User
	Tasks []Task
}
type User struct {
	Uid primitive.ObjectID `json:"id" bson:"uid,omitempty"`
	Name string
	
}

type Task struct {
	ID primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Name string
	Deadline primitive.DateTime
	Budget int
	Description string
	Subtasks []Subtask
	
}
type Comment struct {
	User User
	Text string
	Timestamp primitive.DateTime
}
type Subtask struct {
	ID primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Name string
	Description string 
	Budget int
	Assigned_users[] struct{
		User User
		HasCompleted int
	}
	Updates [] struct{
		User User
		Text string
		Timestamp primitive.DateTime
		
	}
}