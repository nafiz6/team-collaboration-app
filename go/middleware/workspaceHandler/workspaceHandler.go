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


func GetAllTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	
	db.DoNothing();
	
	json.NewEncoder(w).Encode(params["workspace-id"])
	

}

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

func AssignUserToProject(w http.ResponseWriter, r *http.Request) {
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
	

	
	id := params["project-id"]

	fmt.Printf("id: %+v", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("objID: %+v", objID)

	

	insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{  "_id", objID    },
		}, bson.D{
			{ "$push", bson.D{{ "workspaces.$[workspace].users", user }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{bson.D{{"workspace.name", "General"  } }},
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

func EditTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
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
			{  "workspaces.tasks._id", objID    },
		}, bson.D{
			{ "$set", bson.D{{ "workspaces.$[workspace].tasks.$[task].name", newTask.Name }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{bson.D{{"task._id", objID  } },
			bson.D{{ "workspace.tasks", bson.D{{ "$exists", true }}  }},
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
			{  "workspaces.tasks._id", objID    },
		}, bson.D{
			{ "$set", bson.D{{ "workspaces.$[workspace].tasks.$[task].budget", newTask.Budget }},  },
		}, options.Update().SetArrayFilters(options.ArrayFilters{ 
			Filters: []interface{}{bson.D{{"task._id", objID  } },
			bson.D{{ "workspace.tasks", bson.D{{ "$exists", true }}  }},
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

func SubtaskUpdates(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newUpdate Update
	_ = json.NewDecoder(r.Body).Decode(&newUpdate)
	//insert newTask into db
	newUpdate.Timestamp = primitive.NewDateTimeFromTime(time.Now())


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
	var newProject Project
	_ = json.NewDecoder(r.Body).Decode(&newProject)
	newProject.Workspaces = append(newProject.Workspaces, Workspace{
		Name: "General",
		ID: primitive.NewObjectID(),
		Users: []User{},
		Tasks: []Task{},
	})


	fmt.Printf("new task %+v\n", newProject)

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
