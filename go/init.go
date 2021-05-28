package main
/*
package main2

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "context"
	"time"
    "github.com/graphql-go/graphql"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    // (*w).Write([]byte(`{"message": "hello world"}`))
    fmt.Fprintf(w, "HELLO FROM GO")
    
}

func httpJsonHandler(w http.ResponseWriter, r *http.Request) {
    
    enableCors(&w)
    // decoder := json.NewDecoder(r.Body)


    
}
func connectToDb() {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}


func handleGraphQlQuery(query string) {
    fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	// query := `
	// 	{
	// 		hello
	// 	}
	// `
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    
}

func main() {
    connectToDb();
    
    http.HandleFunc("/", httpJsonHandler)
    fmt.Println("Running Go Server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
*/