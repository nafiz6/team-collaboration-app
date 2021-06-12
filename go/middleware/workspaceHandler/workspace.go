package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"teams/middleware/cors"
	"teams/middleware/db"
	. "teams/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProjectWorkspaces(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received projectID: %+v", params["project-id"])

	var workspaces []NewWorkspace

	projectID, err := primitive.ObjectIDFromHex(params["project-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Workspaces.Find(context.Background(), bson.D{
		{"_project_id", projectID},
	})

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewWorkspace
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		workspaces = append(workspaces, elem)
	}

	json.NewEncoder(w).Encode(workspaces)
}

func GetWorkspaceUsers(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received projectID: %+v", params["workspace-id"])

	// var users []struct {
	// 	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 	Role int                //role is diff in each workspace
	// 	Name string             //name isnt changed much and is retrieved often here
	// }

	var workspace NewWorkspace

	workspaceID, err := primitive.ObjectIDFromHex(params["workspace-id"])
	if err != nil {
		panic(err)
	}

	workspaceDetails := db.Workspaces.FindOne(context.Background(), bson.D{
		{"_id", workspaceID},
	})

	decodeErr := workspaceDetails.Decode(&workspace)

	if decodeErr != nil {
		fmt.Printf("Error: %s", decodeErr.Error())
		json.NewEncoder(w).Encode(decodeErr.Error())
		return
	}

	var users = []struct {
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Role int                //role is diff in each workspace
		Name string             //name isnt changed much and is retrieved often here
	}{}

	users = workspace.Users

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Role < users[j].Role
	})

	json.NewEncoder(w).Encode(users)
}

func AssignUserToWorkspace(w http.ResponseWriter, r *http.Request) {
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

	//find user's name and role from db from uid

	userDetails := db.Users.FindOne(context.TODO(), bson.D{{"_id", userID}})

	// var user User
	user := bson.M{"role": uid.Role}
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
		{"workspaces._id", objID},
	}, bson.D{
		{"$push", bson.D{{"workspaces.$[workspace].users", user}}},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{"workspace._id", objID}}},
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

func SetWorkspaceUserRole(w http.ResponseWriter, r *http.Request) {
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

	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	insertResult, err := db.Workspaces.UpdateOne(context.TODO(), bson.D{
		{"_id", objID},
	}, bson.D{
		{"$set", bson.D{{"users.$[user].role", uid.Role}}},
	},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{{"user._id", userID}}},
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

func AssignUserToWorkspaceNew(w http.ResponseWriter, r *http.Request) {
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

	//find user's name and role from db from uid

	userDetails := db.Users.FindOne(context.TODO(), bson.D{{"_id", userID}})

	// var user User
	user := bson.M{"role": uid.Role}
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

	insertResult, err := db.Workspaces.UpdateOne(context.TODO(), bson.D{
		{"_id", objID},
	}, bson.D{
		{"$push", bson.D{{"users", user}}},
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

func GetAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["project-id"])
}

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
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
		bson.M{"_id": objID},
		bson.D{
			{"$push", bson.D{
				{"workspaces", newWorkspace},
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

func CreateWorkspaceNew(w http.ResponseWriter, r *http.Request) {

	cors.EnableCors(&w)
	params := mux.Vars(r)
	var newWorkspace NewWorkspace
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
	newWorkspace.Project_ID = objID

	insertResult, insertErr := db.Projects.InsertOne(
		context.TODO(),
		newWorkspace,
	)

	if insertErr != nil {
		fmt.Printf("Error: %s", insertErr.Error())
		json.NewEncoder(w).Encode(insertErr.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult)
	json.NewEncoder(w).Encode(newWorkspace.ID)

}

func DeleteWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	workspaceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	//find tasks
	cur, err := db.Tasks.Find(context.Background(), bson.D{
		{"_workspace_id", workspaceID},
	})

	var taskIDs = []primitive.ObjectID{}

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewTask
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		taskIDs = append(taskIDs, elem.ID)
	}

	//find subtasks
	cur, err = db.Subtasks.Find(context.Background(), bson.D{
		{"_task_id", bson.D{
			{"$in", taskIDs},
		}},
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

	//delete tasks
	deleteResult, err = db.Tasks.DeleteMany(context.TODO(), bson.D{
		{"_id", bson.D{
			{"$in", taskIDs},
		}},
	})

	//delete workspace
	deleteResult, err = db.Workspaces.DeleteOne(context.TODO(), bson.D{
		{"_id", workspaceID},
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	json.NewEncoder(w).Encode(deleteResult)

}
