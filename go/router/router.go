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
	// fetches all tasks and their subtasks  || NO NEED
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.GetAllTask).Methods("GET", "OPTIONS")

	// create task
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// create subtask
	router.HandleFunc("/api/subtask/{task-id}", workspaceHandler.CreateSubTask).Methods("POST", "OPTIONS") 

	// assign user to workspace
	router.HandleFunc("/api/assign-workspace/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// assign user to projects
	router.HandleFunc("/api/assign-projects/{project-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// assign user to task
	router.HandleFunc("/api/assign-subtask/{subTask-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// edits (renames/ change deadline/ budget)
	router.HandleFunc("/api/update-task/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// mark as completed
	router.HandleFunc("/api/mark-subTask-complete/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// add updates
	router.HandleFunc("/api/subtask-updates/{subTask-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS") // FIZZ



	/*
		WORKSPACE
	*/
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.GetAllWorkspaces).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.CreateWorkspace).Methods("POST", "OPTIONS") // FIZZ




	/*
		PROJECT
	*/
	router.HandleFunc("/api/project", workspaceHandler.GetAllProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/project", workspaceHandler.CreateProject).Methods("POST", "OPTIONS") // FIZZ

	router.HandleFunc("/api/user", workspaceHandler.GetOneUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", workspaceHandler.GetAllUsers).Methods("GET", "OPTIONS")
	return router
}