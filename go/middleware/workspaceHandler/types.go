package workspaceHandler

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
    ID string
    Name string
    Workspaces []Workspace
}
type Workspace struct {
	Name string
	ID primitive.ObjectID
	Users []User
	Tasks []Task
}
type User struct {
	ID primitive.ObjectID
	Name string
	
}

type Task struct {
	ID primitive.ObjectID
	Name string
	Deadline primitive.DateTime
	Budget int
	Comments []Comment
	Subtasks []Subtask
	
}
type Comment struct {
	User User
	Text string
	Timestamp primitive.DateTime
}
type Subtask struct {
	ID primitive.ObjectID
	Name string
	Description string 
	Budget int
	Assigned_users []User
	Updates [] struct{
		user User
		text string
		status int

		
	}
}