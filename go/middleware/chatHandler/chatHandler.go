package chatHandler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"teams/middleware/accountsHandler"
	"teams/middleware/cors"
	. "teams/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

// define our WebSocket endpoint
func ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CHAT CLIENT REQ")
	fmt.Println(r.Host)
	UID := accountsHandler.GetUserId(r)
	fmt.Println("UID: " + UID)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	client := &Client{
		ID:    UID,
		Conn:  ws,
		Pools: make(map[string]*Pool),
	}

	client.Read()
}

func getFileType(filename string) string {
	filenameArray := strings.Split(filename, ".")
	return filenameArray[len(filenameArray)-1]

}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	fmt.Println("File Upload Endpoint Hit")

	// 99MB max
	r.ParseMultipartForm(99 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
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

	fmt.Fprintf(w, "http://localhost:8080/static/"+id)

}
