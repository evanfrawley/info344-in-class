package main

import (
	"testing"
	"fmt"
	"net/http"
	"encoding/json"
	"net/rpc"
	"log"
)

//run these benchmarks using:
// go test -bench=. -benchmem

const urlToSummarize = "http://example.com"

func BenchmarkRPC(b *testing.B) {
	//TODO: implement a benchmark for the RPC server
	client, err := rpc.Dial("tcp", rpcAddr)
	if err != nil {
		b.Fatalf("err dialing rpc server: %v\n", err)
	}
	defer client.Close()

	psum := &PageSummary{}
	for i := 0; i < b.N; i++ {
		if err = client.Call("SummaryService.GetPageSummary", urlToSummarize, psum); err != nil {
			log.Fatalf("error calling RPC %v\n", err)
			if psum.URL != urlToSummarize {
				b.Fatalf("incorrect URL returned. Expected %s but got %s\n", urlToSummarize, psum.URL)
			}
		}
	}
}

func BenchmarkHTTP(b *testing.B) {
	//TODO: implement a benchmark for the HTTP server
	summaryURL := fmt.Sprintf("http://%s?url=%s", httpAddr, urlToSummarize)
	psum := &PageSummary{}
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(summaryURL)
		if err != nil {
			b.Fatalf("err getting page summary: %v\n", err)
		}
		if err = json.NewDecoder(resp.Body).Decode(psum); err != nil {
			b.Fatalf("err decoding page summary: %v\n", err)
		}
		if psum.URL != urlToSummarize {
			b.Fatalf("incorrect URL returned. Expected %s but got %s\n", urlToSummarize, psum.URL)
		}
		resp.Body.Close()
	}
}
