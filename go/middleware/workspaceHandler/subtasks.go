package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teams/middleware/accountsHandler"
	"teams/middleware/cors"
	"teams/middleware/db"
	. "teams/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSubtasks(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received taskID: %+v", params["task-id"])

	var subtasks []NewSubtask

	TaskID, err := primitive.ObjectIDFromHex(params["task-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Subtasks.Find(context.Background(), bson.D{
		{"_task_id", TaskID},
	})

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewSubtask
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		subtasks = append(subtasks, elem)
	}

	json.NewEncoder(w).Encode(subtasks)
}

func GetSubtaskUpdates(w http.ResponseWriter, r *http.Request) {

	type Update struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		UserID    primitive.ObjectID `json:"user_id" bson:"_user_id,omitempty"`
		SubtaskID primitive.ObjectID `json:"subtask_id" bson:"_subtask_id,omitempty"`
		Text      string
		Timestamp primitive.DateTime
		User      UserDetailsNew
	}

	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received subtaskID: %+v", params["subTask-id"])

	var updates []Update

	subtaskID, err := primitive.ObjectIDFromHex(params["subTask-id"])
	if err != nil {
		panic(err)
	}

	// cur, err := db.SubtaskUpdates.Find(context.Background(), bson.D{
	// 	{"_subtask_id", subtaskID},
	// })

	cur, err := db.SubtaskUpdates.Aggregate(context.Background(),
		mongo.Pipeline{
			bson.D{
				{"$match", bson.D{
					{"_subtask_id", subtaskID},
				}},
			},
			bson.D{
				{"$lookup", bson.D{
					{"from", "users"},
					{"localField", "_user_id"},
					{"foreignField", "_id"},
					{"as", "user"},
				}},
			},
			bson.D{
				{"$unwind", "$user"},
			},
		},
	)

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem Update
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		updates = append(updates, elem)
	}

	json.NewEncoder(w).Encode(updates)
}

func AssignUserToSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var uid struct {
		Uid  string
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
	user := bson.M{"has_completed": 0}
	decodeErr := userDetails.Decode(&user)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("user deets %+v", user)

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
		{"workspaces.tasks.subtasks._id", objID},
	}, bson.D{
		{"$push", bson.D{{"workspaces.$[workspace].tasks.$[task].subtasks.$[subtask].assigned_users", user}}},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{{"$and", bson.A{bson.D{{"subtask._id", objID}, {"subtask.assigned_users", bson.D{{
				"$exists", true,
			}}}}}}},
			bson.D{{"workspace.tasks", bson.D{{
				"$exists", true,
			}}}},
			bson.D{{"task.subtasks", bson.D{{
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

func AssignUserToSubTaskNew(w http.ResponseWriter, r *http.Request) {
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

	id := params["subTask-id"]

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
	user := bson.M{"has_completed": 0}
	decodeErr := userDetails.Decode(&user)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	fmt.Printf("user deets %+v", user)

	//check if user is already assigned to subtask
	var subtask NewSubtask

	err = db.Subtasks.FindOne(context.Background(), bson.D{
		{"_id", objID},
		{"assigned_users._id", userID},
	}).Decode(&subtask)

	if err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User already assigned to subtask")
		return
	}

	insertResult, err := db.Subtasks.UpdateOne(context.TODO(), bson.D{
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

func CompleteSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var completion struct {
		Uid    string
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
		{"workspaces.tasks.subtasks._id", objID},
	}, bson.D{
		{"$set", bson.D{{"workspaces.$[].tasks.$[].subtasks.$[subtask].assigned_users.$[user].has_completed", completion.Status}}},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{{"subtask._id", objID}},
			bson.D{{"user._id", userID}},
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

func EditSubtask(w http.ResponseWriter, r *http.Request) {

	//enable edits for name, description, spent
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newSubtask NewSubtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubtask)
	//insert newTask into db

	fmt.Printf("new task %+v\n", newSubtask)

	id := params["subTask-id"]

	fmt.Printf("id: %+v", id)

	subtaskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", subtaskID)

	if newSubtask.Name != "" {
		//name isnt empty
		insertResult, err := db.Subtasks.UpdateOne(context.TODO(), bson.D{
			{"_id", subtaskID},
		}, bson.D{
			{"$set", bson.D{{"name", newSubtask.Name}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newSubtask.Description != "" {
		//description isnt empty
		insertResult, err := db.Subtasks.UpdateOne(context.TODO(), bson.D{
			{"_id", subtaskID},
		}, bson.D{
			{"$set", bson.D{{"description", newSubtask.Description}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)

	}

	if newSubtask.Spent > 0 {
		//budget isnt empty
		insertResult, err := db.Subtasks.UpdateOne(context.TODO(), bson.D{
			{"_id", subtaskID},
		}, bson.D{
			{"$set", bson.D{{"spent", newSubtask.Spent}}},
		})

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(insertResult)
	}

	json.NewEncoder(w).Encode("Successfully updated")
}

func CompleteSubTaskNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var completion struct {
		Uid    string
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

	//check if user is assigned to subtask
	var subtask NewSubtask

	err = db.Subtasks.FindOne(context.Background(), bson.D{
		{"_id", objID},
		{"assigned_users._id", userID},
	}).Decode(&subtask)

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User isnt assigned to subtask")
		return
	}

	insertResult, err := db.Subtasks.UpdateOne(context.TODO(), bson.D{
		{"_id", objID},
	}, bson.D{
		{"$set", bson.D{{"assigned_users.$[user].has_completed", completion.Status}}},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{{"user._id", userID}},
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
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newSubTask Subtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubTask)
	//insert newTask into db

	if newSubTask.Name == "" {

		json.NewEncoder(w).Encode("Name Cannot Be Empty")
		return
	}

	fmt.Printf("new sub task %+v\n", newSubTask)

	newSubTask.ID = primitive.NewObjectID()

	id := params["task-id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v\n", objID)

	findResult := db.Projects.FindOne(context.TODO(), bson.D{{"workspaces.tasks._id", objID}})

	// findResult := db.Projects.FindOne(context.TODO(), bson.D{{ "name", "Cyberpunk" }})

	var elem Project
	errr := findResult.Decode(&elem)
	if errr != nil {
		log.Fatal(err)
	}

	fmt.Printf("found: %+v\n\n", elem)

	// bson.D{{  "workspaces", bson.D{{"$elemMatch", bson.D{{ "tasks", bson.D{{ "$elemMatch", bson.D{{ "_id", objID }} }} }} }}    }}
	insertResult := db.Projects.FindOneAndUpdate(context.TODO(), bson.D{
		{"workspaces.tasks._id", objID},
	}, bson.D{
		{"$push", bson.D{
			{"workspaces.$[workspace].tasks.$[task].subtasks", newSubTask},
		}}}, options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{{"task._id", objID}},
			bson.D{{"workspace.tasks", bson.D{{"$exists", true}}}},
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

func CreateSubTaskNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	params := mux.Vars(r)
	var newSubTask NewSubtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubTask)
	//insert newTask into db

	var self = accountsHandler.GetUserId(r)

	selfID, err := primitive.ObjectIDFromHex(self)
	if err != nil {
		panic(err)
	}

	//get self details

	var userDetails UserDetailsNew

	err = db.Users.FindOne(context.TODO(), bson.D{{"_id", selfID}}).Decode(&userDetails)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if newSubTask.Name == "" {

		json.NewEncoder(w).Encode("Name Cannot Be Empty")
		return
	}

	fmt.Printf("new sub task %+v\n", newSubTask)

	newSubTask.ID = primitive.NewObjectID()
	newSubTask.Date_created = primitive.NewDateTimeFromTime(time.Now())

	id := params["task-id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v\n", objID)

	newSubTask.TaskID = objID

	newSubTask.Assigned_users = []struct {
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name string
		Has_Completed int
	}{
		{
			ID:   userDetails.ID,
			Name: userDetails.Name,
			Has_Completed: 0,
		},
	}

	insertResult, err := db.Subtasks.InsertOne(context.TODO(), newSubTask)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult.InsertedID)
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

func SubtaskUpdates(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newUpdate Update
	_ = json.NewDecoder(r.Body).Decode(&newUpdate)
	//insert newTask into db
	newUpdate.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	if newUpdate.Text == "" {

		json.NewEncoder(w).Encode("Text Cannot Be Empty")
		return
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
		bson.M{"workspaces.tasks.subtasks._id": objID},
		bson.D{
			{"$push", bson.D{
				{"workspaces.$[workspace].tasks.$[task].subtasks.$[subtask].updates", newUpdate},
			}},
		}, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{{"subtask._id", objID}},
				bson.D{{"task.subtasks", bson.D{{"$exists", true}}}},
				bson.D{{"workspace.tasks.subtasks", bson.D{{"$exists", true}}}},
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

func SubtaskUpdatesNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	params := mux.Vars(r)
	var newUpdate NewSubtaskUpdate

	var updateInfo struct {
		Text   string
	}
	_ = json.NewDecoder(r.Body).Decode(&updateInfo)
	//insert newTask into db

	if updateInfo.Text == "" {

		json.NewEncoder(w).Encode("Text Cannot Be Empty")
		return
	}

	id := params["subTask-id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// userID, err := primitive.ObjectIDFromHex(updateInfo.UserID)
	// if err != nil {
	// 	panic(err)
	// }

	var self = accountsHandler.GetUserId(r);

	selfID, err := primitive.ObjectIDFromHex(self)
	if err != nil {
		panic(err)
	}
	newUpdate.Text = updateInfo.Text
	newUpdate.UserID = selfID
	newUpdate.SubtaskID = objID
	newUpdate.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	newUpdate.ID = primitive.NewObjectID()

	fmt.Printf("new sub task update %+v\n", newUpdate)

	fmt.Printf("objID: %+v\n", objID)

	//check if user is assigned to subtask first
	var task NewSubtask

	err = db.Subtasks.FindOne(context.Background(), bson.D{
		{"_id", objID},
		{"assigned_users._id", selfID},
	}).Decode(&task)

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User isn't assigned to subtask")
		return
	}

	insertResult, insertErr := db.SubtaskUpdates.InsertOne(
		context.TODO(), newUpdate)

	// {"workspace.tasks", bson.D{{ "$exists", true }} }

	if insertErr != nil {
		fmt.Printf("Error: %s", insertErr.Error())
		json.NewEncoder(w).Encode(insertErr.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult)
	json.NewEncoder(w).Encode(insertResult)

}

func DeleteSubtaskHelper(subtaskID primitive.ObjectID) (*mongo.DeleteResult, error) {

	deleteResult, err := db.Subtasks.DeleteOne(context.TODO(), bson.D{
		{"_id", subtaskID},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		// json.NewEncoder(w).Encode(err.Error())
		return deleteResult, err
	}

	deleteResult, err = db.SubtaskUpdates.DeleteMany(context.TODO(), bson.D{
		{"_subtask_id", subtaskID},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		// json.NewEncoder(w).Encode(err.Error())
		return deleteResult, err
	}
	return deleteResult, err

}

func DeteleSubtask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// fmt.Printf("body %+v\n", r.Body)
	// _ = json.NewDecoder(r.Body).Decode(&completion)
	//insert newTask into db

	// fmt.Printf("new task %+v\n", completion)

	id := params["subTask-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	deleteResult, err := DeleteSubtaskHelper(objID)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(deleteResult)

	// deleteResult, err := db.Subtasks.DeleteOne(context.TODO(), bson.D{
	// 	{"_id", objID},
	// })
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }

	// deleteResult, err = db.SubtaskUpdates.DeleteMany(context.TODO(), bson.D{
	// 	{"_subtask_id", objID},
	// })
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }

	// json.NewEncoder(w).Encode(deleteResult)

}
