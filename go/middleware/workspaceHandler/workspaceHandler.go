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
			{  "workspaces._id", objID    },
		}, bson.D{
			{ "$push", bson.D{{ "workspaces.$[workspace].tasks", newTask }},  },
		}, options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{bson.D{{"workspace._id", objID  } }},
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

func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newSubTask Subtask
	// fmt.Printf("body %+v\n", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&newSubTask)
	//insert newTask into db


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

func GetAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["project-id"])
}

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newWorkspace Workspace
	_ = json.NewDecoder(r.Body).Decode(&newWorkspace)


	json.NewEncoder(w).Encode(params["project-id"])
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);

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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);

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
	var newProject Project
	_ = json.NewDecoder(r.Body).Decode(&newProject)

	json.NewEncoder(w).Encode(params)
}
