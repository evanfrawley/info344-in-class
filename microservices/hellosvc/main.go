package main

import (
"fmt"
"log"
"net/http"
"os"
)

type HelloHandler struct {
    serviceAddr string
}

func (hh *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "text/plain")
    fmt.Fprintf(w, "WASSUP %s", hh.serviceAddr)
}

func main() {
    addr := os.Getenv("ADDR")
    if len(addr) == 0 {
        addr = ":80"
    }

    http.Handle("/v1/hello", &HelloHandler{addr})
    log.Printf("server is listening at http://%s...", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
