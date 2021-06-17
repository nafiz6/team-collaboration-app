package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
	. "teams/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	db.DoNothing()

	json.NewEncoder(w).Encode(params["workspace-id"])

}

func GetWorkspaceTasks(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received workspaceID: %+v", params["workspace-id"])

	var tasks []NewTask

	workspaceID, err := primitive.ObjectIDFromHex(params["workspace-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Tasks.Find(context.Background(), bson.D{
		{"_workspace_id", workspaceID},
	})

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewTask
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, elem)
	}

	json.NewEncoder(w).Encode(tasks)
}

func GetTaskUsers(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received TaskID: %+v", params["task-id"])

	// var users [] struct {
	// 	User_ID primitive.ObjectID
	// 	Name string
	// }

	taskID, err := primitive.ObjectIDFromHex(params["task-id"])
	if err != nil {
		panic(err)
	}

	// cur, err := db.Tasks.Aggregate(context.Background(), mongo.Pipeline{
	// 	bson.D{
	// 		{"$match", bson.D{
	// 			{"_id", taskID},
	// 		},
	// 		}},
	// 	bson.D{
	// 		{"$project", bson.D{
	// 			{"user_ID", "$assigned_users._id"},
	// 			{"name", "$assigned_users.name"},
	// 			{"assigned_users", 0},

	// 		},
	// 		},
	// 	},
	// })
	task := db.Tasks.FindOne(context.Background(), bson.D{
		{"_id", taskID},
	})

	var decodedTask NewTask
	_ = task.Decode(&decodedTask)

	// if decodeErr != nil {
	// 	fmt.Printf("Error: %s", decodeErr.Error())
	// 	json.NewEncoder(w).Encode(decodeErr.Error())
	// 	return
	// }

	users := decodedTask.Assigned_users

	// if err = cur.All(context.Background(), &users); err != nil {
	// 	panic(err)
	// }
	fmt.Println(users)

	// for cur.Next(context.Background()) {

	// 	// create a value into which the single document can be decoded
	// 	var elem NewTask
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	tasks = append(tasks, elem)
	// }

	json.NewEncoder(w).Encode(users)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newTask Task
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db

	fmt.Printf("new task %+v\n", newTask)

	newTask.ID = primitive.NewObjectID()
	// newTask.Subtasks = []Subtask

	id := params["task-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	if newTask.Name != "" {
		//name isnt empty
		insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{"workspaces.tasks._id", objID},
		}, bson.D{
			{"$set", bson.D{{"workspaces.$[workspace].tasks.$[task].name", newTask.Name}}},
		}, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{{"task._id", objID}},
				bson.D{{"workspace.tasks", bson.D{{"$exists", true}}}},
			},
		}))

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newTask.Description != "" {
		//description isnt empty
		insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{"workspaces.tasks._id", objID},
		}, bson.D{
			{"$set", bson.D{{"workspaces.$[workspace].tasks.$[task].description", newTask.Description}}},
		}, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{{"task._id", objID}},
				bson.D{{"workspace.tasks", bson.D{{"$exists", true}}}},
			},
		}))

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newTask.Budget > 0 {
		//budget isnt empty
		insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{"workspaces.tasks._id", objID},
		}, bson.D{
			{"$set", bson.D{{"workspaces.$[workspace].tasks.$[task].budget", newTask.Budget}}},
		}, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{{"task._id", objID}},
				bson.D{{"workspace.tasks", bson.D{{"$exists", true}}}},
			},
		}))

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)
	}

	json.NewEncoder(w).Encode(newTask.ID.Hex())
}

func EditTaskNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newTask NewTask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db

	fmt.Printf("new task %+v\n", newTask)

	newTask.ID = primitive.NewObjectID()
	// newTask.Subtasks = []Subtask

	id := params["task-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	if newTask.Name != "" {
		//name isnt empty
		insertResult, err := db.Tasks.UpdateOne(context.TODO(), bson.D{
			{"_id", objID},
		}, bson.D{
			{"$set", bson.D{{"name", newTask.Name}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newTask.Description != "" {
		//description isnt empty
		insertResult, err := db.Tasks.UpdateOne(context.TODO(), bson.D{
			{"_id", objID},
		}, bson.D{
			{"$set", bson.D{{"description", newTask.Description}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newTask.Budget > 0 {
		//budget isnt empty
		insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{"_id", objID},
		}, bson.D{
			{"$set", bson.D{{"budget", newTask.Budget}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)
	}

	json.NewEncoder(w).Encode(newTask.ID.Hex())
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newTask Task
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db

	fmt.Printf("new task %+v\n", newTask)

	newTask.ID = primitive.NewObjectID()
	// newTask.Subtasks = []Subtask

	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	insertResult := db.Projects.FindOneAndUpdate(context.TODO(), bson.D{
		{"workspaces._id", objID},
	}, bson.D{
		{"$push", bson.D{{"workspaces.$[workspace].tasks", newTask}}},
	}, options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{"workspace._id", objID}}},
	}),
	)

	doc := bson.M{}
	decodeErr := insertResult.Decode(&doc)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", doc)

	json.NewEncoder(w).Encode(newTask.ID.Hex())

}

func AssignUserToTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var uid struct {
		Uid string
		// Role int
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

	id := params["task-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	//find user's name and role from db from uid

	userDetails := db.Users.FindOne(context.TODO(), bson.D{{"_id", userID}})
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
	user := bson.M{}
	decodeErr := userDetails.Decode(&user)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("user deets %+v", user)

	//check if user already assigned to task
	var task NewTask

	err = db.Tasks.FindOne(context.Background(), bson.D{
		{"_id", objID},
		{"assigned_users._id", userID},
	}).Decode(&task)

	if err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User already assigned to task")
		return
	}


	insertResult, err := db.Tasks.UpdateOne(context.TODO(), bson.D{
		{"_id", objID},
	}, bson.D{
		{"$push", bson.D{{"assigned_users", user}}},
	},
	)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// fmt.Printf("Inserted: %+v\n", doc)

	json.NewEncoder(w).Encode(insertResult)

}

func CreateTaskNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newTask NewTask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db

	fmt.Printf("new task %+v\n", newTask)

	newTask.ID = primitive.NewObjectID()
	newTask.Date_created = primitive.NewDateTimeFromTime(time.Now())
	// newTask.Subtasks = []Subtask

	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	newTask.Workspace_ID = objID

	fmt.Printf("objID: %+v", objID)

	insertResult, err := db.Tasks.InsertOne(context.TODO(), newTask)

	// decodeErr := insertResult.Decode(&doc)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult.InsertedID)

	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	id := params["task-id"]

	fmt.Printf("id: %+v", id)

	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	//find subtasks
	cur, err := db.Subtasks.Find(context.Background(), bson.D{
		{"_id", taskID},
	})

	var subtaskIDs = []primitive.ObjectID{}

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewSubtask
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		subtaskIDs = append(subtaskIDs, elem.ID)
	}

	//delete sabtaskUpdates
	deleteResult, err := db.SubtaskUpdates.DeleteMany(context.TODO(), bson.D{
		{"_subtask_id", bson.D{
			{"$in", subtaskIDs},
		}},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	//delete subtasks
	deleteResult, err = db.Subtasks.DeleteMany(context.TODO(), bson.D{
		{"_id", bson.D{
			{"$in", subtaskIDs},
		}},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	//delete task
	deleteResult, err = db.Tasks.DeleteOne(context.TODO(), bson.D{
		{"_id", taskID},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	json.NewEncoder(w).Encode(deleteResult)

}
