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

    router.HandleFunc("/api/chat", chatHandler.ServeWs)

	/*
	  TASKS
	 */
	// fetches all tasks and their subtasks
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.GetAllTask).Methods("GET", "OPTIONS")

	// create task
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// create subtask
	router.HandleFunc("/api/subtask/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/assign-subtask/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// edits 
	router.HandleFunc("/api/update-task/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/mark-task-complete/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// add notes on progress
	router.HandleFunc("/api/subtask-progress/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")



	/*
		WORKSPACE
	*/
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.GetAllWorkspaces).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.CreateWorkspace).Methods("POST", "OPTIONS")




	/*
		PROJECT
	*/
	router.HandleFunc("/api/project", workspaceHandler.GetAllProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/project", workspaceHandler.GetAllProjects).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/user", workspaceHandler.GetOneUser).Methods("GET", "OPTIONS")
	return router
}