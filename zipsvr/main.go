package main

import (
    "net/http"
    "log"
    "fmt"
    "runtime"
    "encoding/json"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
    w.Header().Add("Access-Control-Allow-Origin", "*")
    fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
    runtime.GC()
    stats := &runtime.MemStats{}
    runtime.ReadMemStats(stats)
    w.Header().Add("Access-Control-Allow-Origin", "*")
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

func main() {
    //fmt.Println("Hello World!")
    mux := http.NewServeMux()

    mux.HandleFunc("/hello", helloHandler)
    mux.HandleFunc("/meme", memoryHandler)

    fmt.Printf("server is litening at http://localhost:4000")
    log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
