package router

import (
	//"teams/middleware/accountsHandler"
	"teams/middleware/chatHandler"
	"teams/middleware/workspaceHandler"
	"github.com/gorilla/mux"
	"net/http"
	
)


// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

    router.HandleFunc("/api/chat", chatHandler.ServeWs)

	/*
	  TASKS
	 */
	// fetches all tasks and their subtasks  || no need
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.GetAllTask).Methods("GET", "OPTIONS")

	// create task
	router.HandleFunc("/api/task/{workspace-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")	//DONE

	// create subtask
	router.HandleFunc("/api/subtask/{task-id}", workspaceHandler.CreateSubTask).Methods("POST", "OPTIONS")	//DONE

	// assign user to subtask
	router.HandleFunc("/api/assign-subtask/{subTask-id}", workspaceHandler.AssignUserToSubTask).Methods("POST", "OPTIONS")	//DONE, but need to select user from workspace instead of Users

	// edits (renames/ change deadline/ budget)
	router.HandleFunc("/api/update-task/{task-id}", workspaceHandler.EditTask).Methods("POST", "OPTIONS")	//DONE for name and budget only

	// mark as completed
	router.HandleFunc("/api/mark-subTask-complete/{subTask-id}", workspaceHandler.CompleteSubTask).Methods("POST", "OPTIONS")	//CODED, NOT CHECKED

	// add updates
	router.HandleFunc("/api/subtask-updates/{subTask-id}", workspaceHandler.SubtaskUpdates).Methods("POST", "OPTIONS") // FIZZ
	
	// delete tasks
	router.HandleFunc("/api/delete-task/{task-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")	

	// delete subtasks
	router.HandleFunc("/api/delete-subTask/{subTask-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")



	/*
		WORKSPACE
	*/
	// fetches all workspace under this project and their tasks  || no need
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.GetAllWorkspaces).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/workspace/{project-id}", workspaceHandler.CreateWorkspace).Methods("POST", "OPTIONS") // FIZZ -- DONE

	// assign user to workspace
	router.HandleFunc("/api/assign-workspace/{workspace-id}", workspaceHandler.AssignUserToWorkspace).Methods("POST", "OPTIONS")	//DONE 

	// delete workspace
	router.HandleFunc("/api/delete-subTask/{subTask-id}", workspaceHandler.CreateTask).Methods("POST", "OPTIONS")



	/*
		PROJECT
	*/
	router.HandleFunc("/api/project", workspaceHandler.GetAllProjects).Methods("GET", "OPTIONS")    //DONE
	router.HandleFunc("/api/project", workspaceHandler.CreateProject).Methods("POST", "OPTIONS") // FIZZ - DONE

	// assign user to projects
	//this adds users to "General" workspace of specified project, since project doesnt have users[]
	router.HandleFunc("/api/assign-projects/{project-id}", workspaceHandler.AssignUserToProject).Methods("POST", "OPTIONS")	//DONE


	router.HandleFunc("/api/user", workspaceHandler.GetOneUser).Methods("GET", "OPTIONS")	//DONE
	router.HandleFunc("/api/users", workspaceHandler.GetAllUsers).Methods("GET", "OPTIONS")	//DONE


	/*
		File server
	*/
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/api/upload-file/{workspace-id}", chatHandler.UploadFile).Methods("POST", "OPTIONS")


	return router
}