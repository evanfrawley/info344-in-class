package main

import (
    "net/http"
    "log"
    "fmt"
    "runtime"
    "encoding/json"
    "info344-in-class/zipsvr/models"
    "strings"
    "os"
    "info344-in-class/zipsvr/handlers"
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
    addr := os.Getenv("GO_ADDR")
    if len(addr) == 0 {
        addr = ":80"
    }

    zips, err := models.LoadZips("./zips.csv")
    if err != nil {
        log.Fatalf("error loading zips: %v", err)
    }

    log.Printf("loaded %d zips", len(zips))

    cityIndex := models.ZipIndex{}
    for _, z := range zips {
        cityLower := strings.ToLower(z.City)
        cityIndex[cityLower] = append(cityIndex[cityLower], z)
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/hello", helloHandler)
    mux.HandleFunc("/meme", memoryHandler)

    const zipsPath = "/zips/"
    cityHandler := &handlers.CityHandler{
        Index: cityIndex,
        PathPrefix: zipsPath,
    }

    mux.Handle(zipsPath, cityHandler)

    fmt.Printf("server is litening at http://%s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
