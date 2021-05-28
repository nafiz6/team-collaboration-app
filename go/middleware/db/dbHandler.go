package db

import(
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"teams/middleware/url"
    "log"
	"fmt"
	"context"
	"time"
)
var Users *mongo.Collection
var Projects *mongo.Collection





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
}

func DoNothing(){
	fmt.Println()
}