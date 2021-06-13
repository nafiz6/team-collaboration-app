package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string
	Workspaces []Workspace
}
type Workspace struct {
	Name  string
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Users []User
	Tasks []Task
}
type User struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string
}

type UserDetailsNew struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string
	Password string
	Username string
	Date_joined primitive.DateTime
	Dp string
	Bio string
}

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string
	Deadline    primitive.DateTime
	Budget      int
	Description string
	Subtasks    []Subtask
}
type Comment struct { //not used anywhere
	User      User
	Text      string
	Timestamp primitive.DateTime
}
type Subtask struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string
	Description    string
	Budget         int
	Assigned_users []struct {
		ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name          string
		Has_Completed int
	}
	Updates []Update
}

type Update struct {
	User      User
	Text      string
	Timestamp primitive.DateTime
}

type NewProject struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string
	Description  string
	Date_created primitive.DateTime
	Budget       int
	// Workspaces []struct {
	// 	ID primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	// }
}

type NewWorkspace struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Project_ID   primitive.ObjectID `json:"project_id" bson:"_project_id,omitempty"`
	Name         string
	Description  string
	Date_created primitive.DateTime
	Users        []struct {
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Role int                //role is diff in each workspace
		Name string             //name isnt changed much and is retrieved often here
	}
}

type NewTask struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Workspace_ID   primitive.ObjectID `json:"workspace_id" bson:"_workspace_id,omitempty"`
	Name           string
	Deadline       primitive.DateTime
	Budget         int
	Description    string
	Spent          int
	Assigned_users []struct { //this represents superset of all subtask users
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name string             //maybe
	}
}

type NewSubtask struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TaskID         primitive.ObjectID `json:"task_id" bson:"_task_id,omitempty"`
	Name           string
	Description    string
	Budget         int
	Assigned_users []struct {
		ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name          string             //maybe
		Has_Completed int
	}
}

type NewSubtaskUpdate struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"_user_id,omitempty"`
	SubtaskID primitive.ObjectID `json:"subtask_id" bson:"_subtask_id,omitempty"`
	Text      string
	Timestamp primitive.DateTime
}
