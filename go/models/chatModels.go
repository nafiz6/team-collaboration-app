package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/tools/go/types/objectpath"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"teams/middleware/cors"
	"teams/middleware/db"
)

var AllPools = make(map[string]*Pool)

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
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: "Connection", Body: "New User Connected..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "Connection", Body: "User Disconnected..."})
			}
			break
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
	Type        string `json:"type"`
	Body        string `json:"body"`
	WorkspaceId string `json:"workspaceId"`
	ClientId    string
}

type DbMessage struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"type"`
	Body        string             `json:"body"`
	WorkspaceId primitive.ObjectID `json:"id" bson:"_workspace_id,omitempty"`
	ClientId    primitive.ObjectID `json:"id" bson:"_user_id,omitempty"`
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

		pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

func createPool(workspaceId string) *Pool {
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
	})

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
