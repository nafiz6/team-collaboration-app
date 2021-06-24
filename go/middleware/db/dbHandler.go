package db

import (
	"context"
	"fmt"
	"log"
	"teams/middleware/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users *mongo.Collection
var Projects *mongo.Collection
var Workspaces *mongo.Collection
var Tasks *mongo.Collection
var Subtasks *mongo.Collection
var SubtaskUpdates *mongo.Collection
var Chats *mongo.Collection
var SubtaskFiles *mongo.Collection

// don't change function name
// called when this package is initialized
func init() {

	client, err := mongo.NewClient(options.Client().ApplyURI(url.Mongo))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	Users = client.Database("collab").Collection("users")
	Projects = client.Database("collab").Collection("projects")
	Workspaces = client.Database("collab").Collection("workspaces")
	Tasks = client.Database("collab").Collection("tasks")
	Subtasks = client.Database("collab").Collection("subtasks")
	SubtaskUpdates = client.Database("collab").Collection("subtask_updates")
	Chats = client.Database("collab").Collection("chats")
	SubtaskFiles = client.Database("collab").Collection("subtask_files")
}

func DoNothing() {
	fmt.Println()
}
