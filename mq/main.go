package main

import (
    "github.com/streadway/amqp"
    "os"
    "fmt"
    "log"
    "encoding/json"
)

type Msg struct {
    User string `json:"user"`
    Password string `json:"password"`
}

func listen(msgs <-chan amqp.Delivery) {
    log.Println("listening for new msgs...")
    for msg := range msgs {
        parsedMsg := &Msg{}
        err := json.Unmarshal(msg.Body, parsedMsg)
        if err != nil {
            fmt.Printf("ran into err: %v", err)
        }
        log.Printf("%+v\n", parsedMsg)
    }
}

func main() {
    mqAddr := os.Getenv("MQADDR")
    if len(mqAddr) == 0 {
        mqAddr = "localhost:5672"
    }

    mqURL := fmt.Sprintf("amqp://%s", mqAddr)

    conn, err := amqp.Dial(mqURL)
    if err != nil {
        log.Fatalf("err connecting to rabbitmq: %v", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf("err creating channel: %v", err)
    }

    q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)

    msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)

    go listen(msgs)

    neverEnd := make(chan bool)
    <-neverEnd
}