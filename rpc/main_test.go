package main

import (
	"testing"
)

//run these benchmarks using:
// go test -bench=. -benchmem

const urlToSummarize = "http://example.com"

func BenchmarkRPC(b *testing.B) {
	//TODO: implement a benchmark for the RPC server
}

func BenchmarkHTTP(b *testing.B) {
	//TODO: implement a benchmark for the HTTP server
}
