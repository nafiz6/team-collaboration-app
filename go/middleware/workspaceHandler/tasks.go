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


func GetAllTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	
	db.DoNothing();
	
	json.NewEncoder(w).Encode(params["workspace-id"])
	

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
	
	if newTask.Description != "" {
		//description isnt empty
		insertResult, err := db.Projects.UpdateOne(context.TODO(), bson.D{
			{  "workspaces.tasks._id", objID    },
		}, bson.D{
			{ "$set", bson.D{{ "workspaces.$[workspace].tasks.$[task].description", newTask.Description }},  },
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

