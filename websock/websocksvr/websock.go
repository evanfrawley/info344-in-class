package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
    "fmt"
)

/*
TODO: Implement the code in this file, according to the comments.
If you haven't yet read the assigned reading, now would be a
good time to do so:
- Read the Overview section of the Gorilla WebSockets package
https://godoc.org/github.com/gorilla/websocket
- Read the Writing WebSocket Client Application
https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications
*/

//WebSocketsHandler is a handler for WebSocket upgrade requests
const WSBufferSize = 1024

type WebSocketsHandler struct {
    notifier *Notifier
    upgrader *websocket.Upgrader
}

//NewWebSocketsHandler constructs a new WebSocketsHandler
func NewWebSocketsHandler(notifer *Notifier) *WebSocketsHandler {
    upgrader := &websocket.Upgrader{
        CheckOrigin:     func(r *http.Request) bool { return true },
        ReadBufferSize:  WSBufferSize,
        WriteBufferSize: WSBufferSize,
    }
    ws := WebSocketsHandler{
        notifier: notifer,
        upgrader: upgrader,
    }
    return &ws
}

//ServeHTTP implements the http.Handler interface for the WebSocketsHandler
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("received websocket upgrade request")
    //TODO: Upgrade the connection to a WebSocket, and add the
    //new websock.Conn to the Notifier. See
    //https://godoc.org/github.com/gorilla/websocket#hdr-Overview
    conn, err := wsh.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Printf("client is getting added\n")
    wsh.notifier.AddClient(conn)
}

//Notifier is an object that handles WebSocket notifications
type Notifier struct {
    clients         []*websocket.Conn
    eventQ          chan []byte
    clientsToAdd    chan *websocket.Conn
    clientsToRemove chan *websocket.Conn
}

//NewNotifier constructs a new Notifier
func NewNotifier() *Notifier {
    n := &Notifier{
        clients:         []*websocket.Conn{},
        eventQ:          make(chan []byte),
        clientsToAdd:    make(chan *websocket.Conn),
        clientsToRemove: make(chan *websocket.Conn),
    }
    go n.start()
    return n
}

//AddClient adds a new client to the Notifier
func (n *Notifier) AddClient(client *websocket.Conn) {
    log.Println("adding new WebSockets client")
    n.clientsToAdd <- client
    for {
        if _, _, err := client.NextReader(); err != nil {
            n.removeClient(client)
            client.Close()
            break
        }
    }
}

//Notify broadcasts the event to all WebSocket clients
func (n *Notifier) Notify(event []byte) {
    log.Printf("adding event to the queue")
    n.eventQ <- event
}

//start starts the notification loop
func (n *Notifier) start() {
    for {
        log.Println("starting notifier loop")
        select {
        case clientToAdd := <-n.clientsToAdd:
            n.clients = append(n.clients, clientToAdd)
        case clientToRemove := <-n.clientsToRemove:
            for i, clientAtIndex := range n.clients {
                if clientAtIndex == clientToRemove {
                    fmt.Printf("client is getting removed from clients \n")
                    n.clients = append(n.clients[:i], n.clients[i+1:]...)
                    break
                }
            }

        case event := <-n.eventQ:
            fmt.Printf("dispatching event: %s\n", string(event))
            for _, client := range n.clients {
                fmt.Printf("writing msg\n")
                if err := client.WriteMessage(websocket.TextMessage, event); err != nil {
                    fmt.Printf("recieved err: %v\n", err)
                    n.clientsToRemove <- client
                }
            }
        }
    }
}

func (n *Notifier) removeClient(client *websocket.Conn) {
    n.clientsToRemove <- client
}
