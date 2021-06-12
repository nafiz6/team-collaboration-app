package main

import (
	"fmt"
	"log"
	"net/http"
	"teams/router"
)

func main() {
	r := router.Router()
	fmt.Println("Running Go Server at port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
