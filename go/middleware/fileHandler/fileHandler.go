package fileHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"teams/middleware/db"
	"time"

	"teams/middleware/cors"
	. "teams/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getFileType(filename string) string {
	filenameArray := strings.Split(filename, ".")
	return filenameArray[len(filenameArray)-1]

}

func TaskGetFile(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	params := mux.Vars(r)

	id := params["task-id"]
	log.Printf("RECEIVED TASK ID" + id)

	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	filesDetails := make([]TaskFile, 0)
	cur, err := db.SubtaskFiles.Find(context.Background(), bson.D{
		{"taskid", taskId},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for cur.Next(context.Background()) {
		var fileDetails TaskFile
		err = cur.Decode(&fileDetails)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return

		}
		filesDetails = append(filesDetails, fileDetails)

	}
	cur.Close(context.Background())
	json.NewEncoder(w).Encode(filesDetails)

}

func WorkspaceGetFiles(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	params := mux.Vars(r)

	id := params["workspace-id"]
	log.Printf("RECEIVED WORKSPACE ID" + id)

	workspaceId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	var filesDetails []WorkspaceFile = []WorkspaceFile{}

	cur, err := db.SubtaskFiles.Find(context.Background(), bson.D{
		{"workspaceid", workspaceId},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for cur.Next(context.Background()) {
		var fileDetails WorkspaceFile
		err = cur.Decode(&fileDetails)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return

		}
		filesDetails = append(filesDetails, fileDetails)

	}
	cur.Close(context.Background())
	json.NewEncoder(w).Encode(filesDetails)
}

func TaskUploadFile(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	fileUrl := UploadFile(w, r)

	var fileDetails TaskFile
	fileDetails.ID = primitive.NewObjectID()
	fileDetails.FileName = r.FormValue("filename")
	fileDetails.Url = fileUrl
	fileDetails.Date_created = primitive.NewDateTimeFromTime(time.Now())
	log.Println("RECEIVED TASK " + r.FormValue("taskId"))
	var err error
	fileDetails.TaskId, err = primitive.ObjectIDFromHex(r.FormValue("taskId"))

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	if fileDetails.TaskId.String() == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("taskid Cannot be empty"))
		return

	}
	if fileDetails.FileName == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Filename Cannot be empty"))
		return

	}

	insertResult, err := db.SubtaskFiles.InsertOne(context.TODO(),
		fileDetails)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult.InsertedID)

	fmt.Fprintf(w, fileUrl)

}

func WorkspaceUploadFile(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	fileUrl := UploadFile(w, r)

	var fileDetails WorkspaceFile
	fileDetails.ID = primitive.NewObjectID()
	fileDetails.FileName = r.FormValue("filename")
	fileDetails.Url = fileUrl
	fileDetails.Date_created = primitive.NewDateTimeFromTime(time.Now())
	var err error
	fileDetails.WorkspaceId, err = primitive.ObjectIDFromHex(r.FormValue("workspaceId"))

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	if fileDetails.WorkspaceId.String() == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Workspace Cannot be empty"))
		return

	}
	if fileDetails.FileName == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Filename Cannot be empty"))
		return

	}

	insertResult, err := db.SubtaskFiles.InsertOne(context.TODO(),
		fileDetails)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error: %s", err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Printf("Inserted: %+v\n", insertResult.InsertedID)
	fmt.Fprintf(w, fileUrl)

}

func BasicUploadFile(w http.ResponseWriter, r *http.Request) {
	url := UploadFile(w, r)

	fmt.Fprintf(w, url)
}

func UploadFile(w http.ResponseWriter, r *http.Request) string {
	cors.EnableCorsCredentials(&w)
	fmt.Println("File Upload Endpoint Hit")

	// 99MB max
	r.ParseMultipartForm(99 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// generate this id

	id := uuid.NewString() + "." + getFileType(handler.Filename)
	tempFile, err := os.Create(filepath.Join("static", id))
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write byte array
	tempFile.Write(fileBytes)

	url := "http://localhost:8080/static/" + id
	return url

}
