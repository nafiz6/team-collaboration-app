package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
	// . "teams/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

//USER STATS
func GetWorkspaceUserTasks(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received workspaceID: %+v", params["workspace-id"])

	workspaceID, err := primitive.ObjectIDFromHex(params["workspace-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Tasks.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"_workspace_id", workspaceID},
			}},
		},
		bson.D{
			{"$unwind", "$assigned_users"},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", "$assigned_users"},
				{"count_tasks_assigned", bson.D{
					{"$sum", 1},
				}},
			}},
		},
	})

	if err != nil {
		panic(err)
	}

	var userTasks = []struct {
		// ID primitive.ObjectID
		ID struct {
			ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
			Name string             //maybe
		} `json:"id" bson:"_id,omitempty"`

		// name  string
		// tasks []NewTask
		Count_tasks_assigned int
	}{}

	// var userTasks = []bson.M{}

	//cur here should return a single value

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem struct {
			// ID primitive.ObjectID
			ID struct {
				ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
				Name string             //maybe
			} `json:"id" bson:"_id,omitempty"`
			// name  string
			// tasks []NewTask
			Count_tasks_assigned int
		}
		// var elem = bson.M{}
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		userTasks = append(userTasks, elem)
	}

	json.NewEncoder(w).Encode(userTasks)
}

//WORKSPACE STATS

//budget: tested
func GetWorkspaceTotalBudget(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received workspaceID: %+v", params["workspace-id"])

	workspaceID, err := primitive.ObjectIDFromHex(params["workspace-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Tasks.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"_workspace_id", workspaceID},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"total_budget", bson.D{
					{"$sum", "$budget"},
				}},
				{"total_spent", bson.D{
					{"$sum", "$spent"},
				}},
			}},
		},
	})

	var totalWorkspaceBudget struct {
		Total_budget int
		Total_spent int
	}

	//cur here should return a single value

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		// var elem NewSubtask
		err := cur.Decode(&totalWorkspaceBudget)
		if err != nil {
			panic(err)
		}

		// subtasks = append(subtasks, elem)
	}

	json.NewEncoder(w).Encode(totalWorkspaceBudget)
}
