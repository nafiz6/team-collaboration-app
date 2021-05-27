package router

import (
	//"teams/middleware/accountsHandler"
	"teams/middleware/chatHandler"
	"teams/middleware/workspaceHandler"
	"github.com/gorilla/mux"
	
)


// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

    router.HandleFunc("/ws", chatHandler.ServeWs)

	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.GetAllWorkspaces).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.CreateWorkspace).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/project", workspaceHandler.GetAllProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/project", workspaceHandler.CreateProject).Methods("POST", "OPTIONS")
	return router
}