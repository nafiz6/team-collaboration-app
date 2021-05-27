package workspaceHandler

import (
	"teams/middleware/cors"
	"teams/models"
	"github.com/gorilla/mux"
	"teams/middleware/db"

    "net/http"
    "encoding/json"
	
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
