package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "golang.org/x/tools/go/types/objectpath"
	"log"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var AllPools = make(map[string]*Pool)

// Each workspace has a pool
type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
	WorkspaceId string
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		// Connect client to the pool
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: "Connection", Body: "New User Connected..."})
			}
			break
		// Disconnect client
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "Connection", Body: "User Disconnected..."})
			}
			break
		// Send message to all clients in the pool
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Pools map[string]*Pool
}

type Message struct {
	Type        string `json:"Type"`
	Body        string `json:"Body"`
	WorkspaceId string `json:"WorkspaceId"`
	ClientId    string
}

type DbMessage struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"Type"`
	Body        string             `json:"Body"`
	WorkspaceId primitive.ObjectID `json:"WorkspaceId" bson:"_workspace_id,omitempty"`
	ClientId    primitive.ObjectID `json:"ClientId" bson:"_user_id,omitempty"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
}

func (c *Client) Read() {
	defer func() {
		for pool, _ := range c.Pools {
			c.Pools[pool].Unregister <- c

		}
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//message := Message{Type: messageType, Body: string(p), WorkspaceId: "123"}
		message := Message{}
		json.Unmarshal([]byte(p), &message)
		message.ClientId = c.ID

		log.Printf(message.Body)

		var pool = createPool(message.WorkspaceId)
		addClientToPool(c, message.WorkspaceId)

		if message.Type != "Connection" {
			AddChatToDb(message)
		}

		pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

func createPool(workspaceId string) *Pool {
	// If pool exists, return
	if val, ok := AllPools[workspaceId]; ok {
		return val
	}
	var pool = NewPool()
	pool.WorkspaceId = workspaceId
	AllPools[workspaceId] = pool
	go pool.Start()
	return pool

}

func addClientToPool(client *Client, workspaceId string) {

	if _, ok := client.Pools[workspaceId]; ok {
		return
	}
	client.Pools[workspaceId] = AllPools[workspaceId]
	AllPools[workspaceId].Register <- client

}

func AddChatToDb(message Message) {

	var newChat DbMessage

	newChat.ID = primitive.NewObjectID()
	newChat.Body = message.Body

	newChat.Type = message.Type

	newChat.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	clientID, err := primitive.ObjectIDFromHex(message.ClientId)
	if err != nil {
		panic(err)
	}

	workspaceID, err := primitive.ObjectIDFromHex(message.WorkspaceId)
	if err != nil {
		panic(err)
	}

	newChat.ClientId = clientID
	newChat.WorkspaceId = workspaceID
	insertResult, err := db.Chats.InsertOne(context.TODO(), newChat)

	if err != nil {
		panic(err)
	}

	print(insertResult)

	// json.NewEncoder(w).Encode(insertResult.InsertedID)

}

//add limit to this later
func GetChats(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	params := mux.Vars(r)
	// _ = json.NewDecoder(r.Body).Decode(&p)

	fmt.Printf("received taskID: %+v", params["workspace-id"])

	var chats []DbMessage

	workspaceID, err := primitive.ObjectIDFromHex(params["workspace-id"])
	if err != nil {
		panic(err)
	}

	cur, err := db.Chats.Find(context.Background(), bson.D{
		{"_workspace_id", workspaceID},
	}, options.Find().SetSort(
		bson.D{
			{"timetamp", -1},
		},
	))

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem DbMessage
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		chats = append(chats, elem)
	}

	json.NewEncoder(w).Encode(chats)
}
