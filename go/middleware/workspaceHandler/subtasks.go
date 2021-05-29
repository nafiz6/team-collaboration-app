package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
	. "teams/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AssignUserToSubTask(w http.ResponseWriter, r *http.Request) {
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

	id := params["subTask-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	//find user's name and role from db from uid


	userDetails := db.Users.FindOne(context.TODO(), bson.D{{ "_id", userID }})
	//LATER find user from this workspace
	// userDetails := db.Projects.FindOne(context.TODO(),
	// 		bson.D{{
	// 			"workspaces.users._id", userID,
	// 		},
	// 		{
	// 			"workspaces.tasks.subtasks._id", objID,
	// 		}},
	// 	)

	// var user User
	user := bson.M{ "has_completed": 0}
	decodeErr := userDetails.Decode(&user)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("user deets %+v", user)
	

	
	

	

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{  "workspaces.tasks.subtasks._id", objID    },
		}, bson.D{
			{ "$push", bson.D{{ "workspaces.$[workspace].tasks.$[task].subtasks.$[subtask].assigned_users", user }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{
				bson.D{{ "$and", bson.A{ bson.D{{ "subtask._id", objID }, { "subtask.assigned_users", bson.D{{
					"$exists", true,
				}}} } }   } }, 
				bson.D{{ "workspace.tasks", bson.D{{
					"$exists", true,
				}}}},
				bson.D{{ "task.subtasks", bson.D{{
					"$exists", true,
				}}}},
			},
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

func CompleteSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var completion struct {
		Uid string
		Status int
	}
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&completion)
	//insert newTask into db


	fmt.Printf("new task %+v\n", completion)



	
	id := params["subTask-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	userID, err := primitive.ObjectIDFromHex(completion.Uid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("userID: %+v", userID)

	

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{  "workspaces.tasks.subtasks._id", objID    },
		}, bson.D{
			{ "$set", bson.D{{ "workspaces.$[].tasks.$[].subtasks.$[subtask].assigned_users.$[user].has_completed", completion.Status }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{
				bson.D{{"subtask._id", objID  }},
				bson.D{{"user._id", userID }},
			 },
		}),
	)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(insertResult)


}


func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newSubTask Subtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubTask)
	//insert newTask into db

	if (newSubTask.Name == ""){

		json.NewEncoder(w).Encode("Name Cannot Be Empty");
		return;
	}


	fmt.Printf("new sub task %+v\n", newSubTask)

	newSubTask.ID = primitive.NewObjectID()


	
	id := params["task-id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v\n", objID)

	findResult := db.Projects.FindOne(context.TODO(), bson.D{{ "workspaces.tasks._id", objID }})

	// findResult := db.Projects.FindOne(context.TODO(), bson.D{{ "name", "Cyberpunk" }})


	var elem Project
	errr := findResult.Decode(&elem)
	if errr != nil {
		log.Fatal(err)
	}

	fmt.Printf("found: %+v\n\n", elem)

	

	// bson.D{{  "workspaces", bson.D{{"$elemMatch", bson.D{{ "tasks", bson.D{{ "$elemMatch", bson.D{{ "_id", objID }} }} }} }}    }}
	insertResult := db.Projects.FindOneAndUpdate(context.TODO(), bson.D{
		{ "workspaces.tasks._id", objID },
	}, bson.D{
		{ "$push", bson.D{
			{ "workspaces.$[workspace].tasks.$[task].subtasks", newSubTask },
		}  }}, 	options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{
				bson.D{{"task._id", objID  }}, 
				bson.D{{ "workspace.tasks", bson.D{{ "$exists", true }}  }},
			},
		},
	))


	
	// {"workspace.tasks", bson.D{{ "$exists", true }} }




	doc := bson.M{}
	decodeErr := insertResult.Decode(&doc)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}


	fmt.Printf("Inserted: %+v\n", doc)
	json.NewEncoder(w).Encode(newSubTask.ID.Hex())




}

func SubtaskUpdates(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newUpdate Update
	_ = json.NewDecoder(r.Body).Decode(&newUpdate)
	//insert newTask into db
	newUpdate.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	if (newUpdate.Text == ""){

		json.NewEncoder(w).Encode("Text Cannot Be Empty");
		return;
	}

	fmt.Printf("new sub task %+v\n", newUpdate)

	
	id := params["subTask-id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v\n", objID)


	insertResult, insertErr := db.Projects.UpdateOne(
		context.TODO(), 
		bson.M{ "workspaces.tasks.subtasks._id": objID },
	    bson.D{
			{ "$push", bson.D{
				{ "workspaces.$[workspace].tasks.$[task].subtasks.$[subtask].updates", newUpdate },
			}},
		},	options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{
				bson.D{{"subtask._id", objID  }}, 
				bson.D{{ "task.subtasks", bson.D{{ "$exists", true }}  }}, 
				bson.D{{ "workspace.tasks.subtasks", bson.D{{ "$exists", true }}  }},
			},
		},
	))

	// {"workspace.tasks", bson.D{{ "$exists", true }} }

	if insertErr != nil {
		fmt.Printf("Error: %s", insertErr.Error())
		json.NewEncoder(w).Encode(insertErr.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult)
	json.NewEncoder(w).Encode(insertResult)

}
