package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"gopkg.in/mgo.v2"
	"github.com/evanfrawley/info344-in-class/tasksvr/models/tasks"
	"github.com/evanfrawley/info344-in-class/tasksvr/handlers"
)

const defaultAddr = ":80"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	//TODO: make connection to the DBMS
	//construct the appropriate tasks.Store
	//construct the handlers.Context


	mongoSess, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}
	mongoStore := tasks.NewMongoStore(mongoSess, "tasks", "tasks")
	handlerContext := handlers.NewHandlerContext(mongoStore)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks", handlerContext.TasksHandler)
	mux.HandleFunc("/v1/tasks/", handlerContext.SpecificTaskHandler)

	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
