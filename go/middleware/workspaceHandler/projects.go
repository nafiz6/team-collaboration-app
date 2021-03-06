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

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)

	var projects []Project

	cur, err := db.Projects.Find(context.Background(), bson.D{{}})

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem Project
		err := cur.Decode(&elem)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Println(err)
			return
		}

		projects = append(projects, elem)
	}
	// Close the cursor once finished
	cur.Close(context.Background())

	fmt.Printf("Found a single document: %+v\n", projects)
	json.NewEncoder(w).Encode(projects)
}

func GetSingleProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	projectID, err := primitive.ObjectIDFromHex(params["project-id"])
	if err != nil {
		panic(err)
	}

	var project NewProject

	err = db.Projects.FindOne(context.Background(), bson.D{{
		"_id", projectID,
	}}).Decode(&project)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time

	fmt.Printf("Found projects: %+v\n", project)
	json.NewEncoder(w).Encode(project)
}

func GetAllProjectsNew(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	var projects []NewProject = []NewProject{}

	var self = accountsHandler.GetUserId(r)
	log.Print(self)
	selfID, err := primitive.ObjectIDFromHex(self)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid User")
		return
	}

	var projectIDs []primitive.ObjectID = []primitive.ObjectID{}

	cur, err := db.Workspaces.Find(context.Background(), bson.D{
		{"name", "General"},
		{"users._id", selfID},
	})

	log.Print(cur)

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewWorkspace
		err := cur.Decode(&elem)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Println(err)
			return
		}

		projectIDs = append(projectIDs, elem.Project_ID)
	}

	//end here

	cur, err = db.Projects.Find(context.Background(),
		// bson.D{},
		//ONLY GET MY PROJECTS
		bson.D{
			{"_id", bson.D{
				{"$in", projectIDs},
			}},
		},
	)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewProject
		err := cur.Decode(&elem)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Println(err)
			return
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
	cors.EnableCorsCredentials(&w)
	var newProject NewProject
	_ = json.NewDecoder(r.Body).Decode(&newProject)

	//CREATE NEW WORKSPACE GENERAL AND LINK TO THIS PROJECT
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

	var self = accountsHandler.GetUserId(r);

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

	generalWorkspace.Users = []struct {
		ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Role int
		Name string
	}{
		{
			ID:   userDetails.ID,
			Role: 0,
			Name: userDetails.Name,
		},
	}
	//add this user to general workspace, CHECK

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
	cors.EnableCorsCredentials(&w)
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

	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", projectID)

	//only allow assigning if self user role is 0

	var self = accountsHandler.GetUserId(r);

	selfID, err := primitive.ObjectIDFromHex(self)
	if err != nil {
		panic(err)
	}

	//get self details

	var myUserDetails UserDetailsNew

	err = db.Users.FindOne(context.TODO(), bson.D{{"_id", selfID}}).Decode(&myUserDetails)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var workspace NewWorkspace

	err = db.Workspaces.FindOne(context.Background(),
		bson.D{
			{"_project_id", projectID},
			{"name", "General"},
			{"users._id", selfID},
			{"users.role", 0},
		},
	).Decode(&workspace)

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User role too low to assign users to project")
		return
	}

	//check if user exists in project first

	// var workspace NewWorkspace

	err = db.Workspaces.FindOne(context.Background(), bson.D{
		{"_project_id", projectID},
		{"name", "General"},
		{"users._id", userID},
	}).Decode(&workspace)

	if err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("User already in project")
		return
	}

	//find workspaces with name "general" and projectID = id
	insertResult, err := db.Workspaces.UpdateOne(context.TODO(), bson.D{
		{"_project_id", projectID},
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
		log.Println(err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Println(err)
			return
		}

		users = append(users, elem)
	}
	// Close the cursor once finished
	cur.Close(context.Background())

	fmt.Println((users))
	json.NewEncoder(w).Encode((users))

}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)

	params := mux.Vars(r)

	id := params["user-id"]

	fmt.Printf("id: %+v", id)

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	var userDetails UserDetailsNew

	err = db.Users.FindOne(context.Background(), bson.D{
		{"_id", userID},
	}).Decode(&userDetails)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(userDetails)

}

func GetMyUserDetails(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	uid := accountsHandler.GetUserId(r)

	userID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		panic(err)
	}
	var userDetails UserDetailsNew

	err = db.Users.FindOne(context.Background(), bson.D{
		{"_id", userID},
	}).Decode(&userDetails)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(userDetails)

}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
	var user User

	err := db.Users.FindOne(context.Background(), bson.D{{}}).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
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
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())

	}

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
