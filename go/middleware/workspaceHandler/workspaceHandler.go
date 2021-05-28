package workspaceHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
	"teams/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetAllTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	
	db.DoNothing();
	
	json.NewEncoder(w).Encode(params["workspace-id"])
	

}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newTask Task
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db
	json.NewEncoder(w).Encode(params["workspace-id"])


	fmt.Printf("new task %+v\n", newTask)

	newTask.ID = primitive.NewObjectID()


	
	id := params["workspace-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	

	insertResult := db.Projects.FindOneAndUpdate(context.TODO(), bson.D{{  "workspaces", bson.D{{"$elemMatch",bson.D{{"_id", objID}}}}    }}, bson.D{{ "$push", bson.D{{ "workspaces.$[workspace].tasks", newTask }},  }}, options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{ Filters: []interface{}{bson.D{{"workspace._id", bson.D{{ "$eq", objID  }}   } }}}))
	




	doc := bson.M{}
	decodeErr := insertResult.Decode(&doc)

	if decodeErr != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Printf("Inserted: %+v\n", doc)




}

func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newSubTask Subtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubTask)
	//insert newTask into db
	json.NewEncoder(w).Encode(params["task-id"])


	fmt.Printf("new sub task %+v\n", newSubTask)

	newSubTask.ID = primitive.NewObjectID()


	
	id := "60b0aa2b60c6feaec02b30f8"

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	

	insertResult := db.Projects.FindOneAndUpdate(context.TODO(), bson.D{{  "workspaces", bson.D{{"$elemMatch",bson.D{{"_id", objID}}}}    }}, bson.D{{ "$push", bson.D{{ "workspaces.$[workspace].tasks", newSubTask }},  }}, options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{ Filters: []interface{}{bson.D{{"workspace._id", bson.D{{ "$eq", objID  }}   } }}}))
	




	doc := bson.M{}
	decodeErr := insertResult.Decode(&doc)

	if decodeErr != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Printf("Inserted: %+v\n", doc)




}

func GetAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["project-id"])
}

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newWorkspace models.Workspace
	_ = json.NewDecoder(r.Body).Decode(&newWorkspace)


	json.NewEncoder(w).Encode(params["project-id"])
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
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
	


}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
	var users []User

	cur, err := db.Users.Find(context.Background(), bson.D{{}})

	if err != nil {
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

}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
	var user User

	err := db.Users.FindOne(context.Background(), bson.D{{}}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("Found a single document: %+v\n", user)

}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newProject models.Project
	_ = json.NewDecoder(r.Body).Decode(&newProject)

	json.NewEncoder(w).Encode(params)
}
