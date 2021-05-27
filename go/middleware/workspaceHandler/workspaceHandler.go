package workspaceHandler

import (
	"teams/middleware/cors"
	"teams/middleware/url"
	"teams/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"

    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "context"
	"time"
)


// don't change function name
// called when this package is initialized
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(url.Mongo))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}


func GetAllTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["workspace-id"])

}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newTask models.Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	//insert newTask into db





	json.NewEncoder(w).Encode(params["workspace-id"])

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
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w);
	params := mux.Vars(r)
	var newProject models.Project
	_ = json.NewDecoder(r.Body).Decode(&newProject)

	json.NewEncoder(w).Encode(params)
}
