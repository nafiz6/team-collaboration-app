package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
	. "teams/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func AssignUserToWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var uid struct {
		Uid string
		Role int
	}
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&uid)
	//insert newTask into db


	fmt.Printf("received user id %+v\n", uid)

	userID, err := primitive.ObjectIDFromHex(uid.Uid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("userID: %+v", userID)

	//find user's name and role from db from uid


	userDetails := db.Users.FindOne(context.TODO(), bson.D{{ "_id", userID }})

	// var user User
	user := bson.M{ "role": uid.Role}
	decodeErr := userDetails.Decode(&user)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("user deets %+v", user)
	

	
	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{  "workspaces._id", objID    },
		}, bson.D{
			{ "$push", bson.D{{ "workspaces.$[workspace].users", user }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{bson.D{{"workspace._id", objID  } }},
		}),
	)
	


	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// fmt.Printf("Inserted: %+v\n", doc)

	json.NewEncoder(w).Encode(insertResult)



}

func GetAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["project-id"])
}

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newWorkspace Workspace
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newWorkspace)
	//insert newTask into db


	fmt.Printf("new workspace %+v\n", newWorkspace)

	newWorkspace.ID = primitive.NewObjectID()
	// newTask.Subtasks = []Subtask


	
	id := params["project-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	

	insertResult, insertErr := db.Projects.UpdateOne(
		context.TODO(), 
		bson.M{ "_id": objID },
	    bson.D{
			{ "$push", bson.D{
				{ "workspaces", newWorkspace },
			}},
		},
	)

	if insertErr != nil {
		fmt.Printf("Error: %s", insertErr.Error())
		json.NewEncoder(w).Encode(insertErr.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult)
	json.NewEncoder(w).Encode(newWorkspace.ID)




}