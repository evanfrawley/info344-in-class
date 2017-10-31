package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func client(index int, c *Cache) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		key := strconv.Itoa(rand.Intn(10000))
		value := key + " value"

		log.Printf("client %d setting %s=%s", index, key, value)
		c.Set(key, value, time.Second*10)
		if value2, found := c.Get(key); !found || value2 != value {
			log.Fatal("incorrect value returned from cache")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := NewCache()
	for i := 0; i < 10; i++ {
		go client(i, c)
	}

	time.Sleep(time.Minute)
}
