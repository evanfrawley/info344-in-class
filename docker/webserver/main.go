package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const headerContentType = "Content-Type"
const contentTypePlainText = "text/plain"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypePlainText)
	fmt.Fprintf(w, "the current time is %v", time.Now().Format("3:04:05PM"))
}

func main() {
	addr := os.Getenv("ADDR")
	//default address to port 80
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
