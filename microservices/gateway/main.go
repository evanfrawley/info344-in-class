package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"net/http/httputil"
	"sync"
	"encoding/json"
)

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello from the gateway! Try requesting /v1/time")
}

type User struct {
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
}

func GetCurrentUser(r *http.Request) *User {
	return &User {
		FirstName: "Test",
		LastName: "User",
	}
}

func NewServiceProxy(addrs []string) *httputil.ReverseProxy {
	nextIndex := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			//modify request to indicate remote host
			user := GetCurrentUser(r)
			userJSON, err := json.Marshal(user)
			if err != nil {
				log.Printf("error marshalling user: %v\n", err)
			}
			r.Header.Add("X-User", string(userJSON))
			mx.Lock()
			r.URL.Host = addrs[nextIndex % len(addrs)]
			nextIndex++
			r.URL.Scheme = "http"
			mx.Unlock()
		},
	}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	//TODO: get network addresses for our
	//timesvc instances
	//timesvc instances
	timeSvcAddrs := os.Getenv("TIMESVC_ADDRS")
	timeHelloAddrs := os.Getenv("HELLOSVC_ADDRS")
	splitTimeSvcs := strings.Split(timeSvcAddrs, ",")
	splitHelloSvcs := strings.Split(timeHelloAddrs, ",")
	fmt.Printf("%v\n", splitHelloSvcs)

	nodeSvcAddrs := os.Getenv("NODESVC_ADDRS")
	fmt.Printf("%s\n", nodeSvcAddrs)
	splitNodeSvcAddrs := strings.Split(nodeSvcAddrs, ",")


	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	//TODO: add reverse proxy handler for `/v1/time`

	mux.Handle("/v1/time", NewServiceProxy(splitTimeSvcs))
	mux.Handle("/v1/hello", NewServiceProxy(splitHelloSvcs))
	mux.Handle("/v1/users/me/hello", NewServiceProxy(splitNodeSvcAddrs))
	mux.Handle("/v1/channels", NewServiceProxy(splitNodeSvcAddrs))

	log.Printf("server is listening at https://%s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "tls/fullchain.pem", "tls/privkey.pem", mux))
}
