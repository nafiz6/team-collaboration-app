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

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)

	var projects []Project

	cur, err := db.Projects.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem Project
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		projects = append(projects, elem)
	}
	// Close the cursor once finished
	cur.Close(context.Background())

	fmt.Printf("Found a single document: %+v\n", projects)
	json.NewEncoder(w).Encode(projects)
}

func GetAllProjectsNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)

	var projects []NewProject

	cur, err := db.Projects.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewProject
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		projects = append(projects, elem)
	}
	// Close the cursor once finished
	cur.Close(context.Background())

	fmt.Printf("Found projects: %+v\n", projects)
	json.NewEncoder(w).Encode(projects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var newProject Project
	_ = json.NewDecoder(r.Body).Decode(&newProject)
	newProject.Workspaces = append(newProject.Workspaces, Workspace{
		Name:  "General",
		ID:    primitive.NewObjectID(),
		Users: []User{},
		Tasks: []Task{},
	})

	// fmt.Printf("Project %+v\n", newProject)

	// fmt.Printf("new task %+v\n", newProject)

	newProject.ID = primitive.NewObjectID()
	// newTask.Subtasks = []Subtask

	insertResult, err := db.Projects.InsertOne(context.TODO(), newProject)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult.InsertedID)

	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

func CreateProjectNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var newProject NewProject
	_ = json.NewDecoder(r.Body).Decode(&newProject)

	//CREATE NEW WORKSPACE GENERAL AND LINK TO THIS PROJECT

	// newProject.Workspaces = append(newProject.Workspaces, Workspace{
	// 	Name: "General",
	// 	ID: primitive.NewObjectID(),
	// 	Users: []User{},
	// 	Tasks: []Task{},
	// })
	fmt.Printf("%+v", newProject)

	newProject.ID = primitive.NewObjectID()
	newProject.Date_created = primitive.NewDateTimeFromTime(time.Now())

	insertProjectResult, err := db.Projects.InsertOne(context.TODO(), newProject)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var generalWorkspace NewWorkspace

	generalWorkspace.ID = primitive.NewObjectID()
	generalWorkspace.Project_ID = newProject.ID
	generalWorkspace.Name = "General"
	generalWorkspace.Date_created = primitive.NewDateTimeFromTime(time.Now())
	generalWorkspace.Description = "General workspace forms global scope for project"
	generalWorkspace.Users = []struct {
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Role int
		Name string
	}{}
	//add this user to general workspace

	insertWorkspaceResult, err := db.Workspaces.InsertOne(context.TODO(), generalWorkspace)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertWorkspaceResult.InsertedID)

	json.NewEncoder(w).Encode(insertProjectResult.InsertedID)

}

func AssignUserToProject(w http.ResponseWriter, r *http.Request) {
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

	id := params["project-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
		{"_id", objID},
	}, bson.D{
		{"$push", bson.D{{"workspaces.$[workspace].users", user}}},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{"workspace.name", "General"}}},
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

func AssignUserToProjectNew(w http.ResponseWriter, r *http.Request) {
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

	id := params["project-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	//find workspaces with name "general" and projectID = id
	insertResult, err := db.Workspaces.UpdateOne(context.TODO(), bson.D{
		{"_project_id", id},
		{"name", "General"},
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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)

	var users []User

	cur, err := db.Users.Find(context.Background(), bson.D{{}})

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}
	// Close the cursor once finished
	cur.Close(context.Background())

	fmt.Println((users))
	json.NewEncoder(w).Encode((users))

}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
	var user User

	err := db.Users.FindOne(context.Background(), bson.D{{}}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", user)

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	id := params["project-id"]

	fmt.Printf("id: %+v", id)

	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	//find workspaces
	cur, err := db.Workspaces.Find(context.Background(), bson.D{
		{"_project_id", projectID},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	var workspaceIDs = []primitive.ObjectID{}

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewWorkspace
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		workspaceIDs = append(workspaceIDs, elem.ID)
	}
	fmt.Printf("workspace IDS %+v", workspaceIDs)

	//find tasks
	cur, err = db.Tasks.Find(context.Background(), bson.D{
		{"_workspace_id", bson.D{
			{"$in", workspaceIDs},
		}},
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	print("WORKSSS")

	var taskIDs  = []primitive.ObjectID{}

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
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

	var subtaskIDs  = []primitive.ObjectID{}

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

	//delete workspaces
	deleteResult, err = db.Workspaces.DeleteMany(context.TODO(), bson.D{
		{"_id", bson.D{
			{"$in", workspaceIDs},
		}},
	})
	json.NewEncoder(w).Encode(deleteResult)

	//delete project
	deleteResult, err = db.Projects.DeleteOne(context.TODO(), bson.D{
		{"_id", projectID},
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	json.NewEncoder(w).Encode(deleteResult)

}
