package router

import (
	//"teams/middleware/accountsHandler"
	"net/http"
	"teams/middleware/accountsHandler"
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

	router.HandleFunc("/api/task-users/{task-id}", workspaceHandler.GetTaskUsers).Methods("GET", "OPTIONS")
	// fetches all tasks and their subtasks  || no need
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.GetWorkspaceTasks).Methods("GET", "OPTIONS")

	// create task
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.CreateTaskNew).Methods("POST", "OPTIONS")	//DONE

	// create subtask
	router.HandleFunc("/api/subtask/{task-id}", workspaceHandler.CreateSubTaskNew).Methods("POST", "OPTIONS")	//DONE

	// assign user to subtask
	router.HandleFunc("/api/assign-subtask/{subTask-id}", workspaceHandler.AssignUserToSubTaskNew).Methods("POST", "OPTIONS")	//DONE, but need to select user from workspace instead of Users

	// assign user to task
	router.HandleFunc("/api/assign-task/{task-id}", workspaceHandler.AssignUserToTask).Methods("POST", "OPTIONS")	//DONE, but need to select user from workspace instead of Users

	// edits (renames/ change deadline/ budget)
	router.HandleFunc("/api/update-task/{task-id}", workspaceHandler.EditTaskNew).Methods("POST", "OPTIONS")	//DONE for name and budget only

	// mark as completed
	router.HandleFunc("/api/mark-subTask-complete/{subTask-id}", workspaceHandler.CompleteSubTaskNew).Methods("POST", "OPTIONS")	//CODED, NOT CHECKED

	// add updates
	router.HandleFunc("/api/subtask-updates/{subTask-id}", workspaceHandler.SubtaskUpdatesNew).Methods("POST", "OPTIONS") // FIZZ
	
	// delete tasks
	router.HandleFunc("/api/delete-task/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	// delete subtasks
	router.HandleFunc("/api/delete-subTask/{subTask-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")

	/*
		WORKSPACE
	*/
	// fetches all workspace under this project and their tasks  || no need
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.GetProjectWorkspaces).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.CreateWorkspaceNew).Methods("POST", "OPTIONS") // FIZZ -- DONE

	// assign user to workspace
	router.HandleFunc("/api/assign-workspace/{workspace-id}", workspaceHandler.AssignUserToWorkspaceNew).Methods("POST", "OPTIONS")	//DONE 

	// delete workspace
	router.HandleFunc("/api/delete-subTask/{subTask-id}", workspaceHandler.CreateTaskNew).Methods("POST", "OPTIONS")

	/*
		PROJECT
	*/
	router.HandleFunc("/api/project", workspaceHandler.GetAllProjectsNew).Methods("GET", "OPTIONS")    //DONE
	router.HandleFunc("/api/project", workspaceHandler.CreateProjectNew).Methods("POST", "OPTIONS") // FIZZ - DONE

	// assign user to projects
	//this adds users to "General" workspace of specified project, since project doesnt have users[]
	router.HandleFunc("/api/assign-projects/{project-id}", workspaceHandler.AssignUserToProject).Methods("POST", "OPTIONS") //DONE

	router.HandleFunc("/api/user", workspaceHandler.GetOneUser).Methods("GET", "OPTIONS")   //DONE
	router.HandleFunc("/api/users", workspaceHandler.GetAllUsers).Methods("GET", "OPTIONS") //DONE

	/*
		File server
	*/
	fs := http.FileServer(http.Dir("./static/"))
	// serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// upload files
	router.HandleFunc("/api/upload-file/{workspace-id}", chatHandler.UploadFile).Methods("POST", "OPTIONS")

	/*
		Authentication
	*/
	router.HandleFunc("/api/login", accountsHandler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/secret-page", accountsHandler.SecretPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/register", accountsHandler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/logout", accountsHandler.Logout).Methods("GET", "OPTIONS")

	return router
}
