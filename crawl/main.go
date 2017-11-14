package main

import (
	"fmt"
	"os"
)

const usage = `
usage:
	crawl <starting-url>
`

//numWorkers is the number of worker goroutines
//we will start: begin with just 1 and increase
//to see the benefits of concurrent execution,
//but don't increase beyond the number of concurrent
//socket connections allowed by your OS
const numWorkers = 1

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	//use the first argument as our starting URL
	//startingURL := os.Args[1]

	//TODO: build a concurrent web crawler
	//with `numWorkers` worker goroutines,
	//using a channel to pass URLs to fetch
	//form the main goroutine to the workers,
	//and a channel to pass *PageLinks structs
	//from the workers back to the main goroutine.
	//Use the `GetPageLinks()` function in `links.go`
	//from your worker goroutines to fetch links
	//for a given URL.
}
