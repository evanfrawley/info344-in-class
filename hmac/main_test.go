package main

import (
    "testing"
    "io"
    "os"
    "bufio"
)

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
	//r, _ :=

	f, _ := os.Open("somestringtoread.txt")
	cases := []struct{
	    name           string
	    key            string
	    stream         io.Reader
	    expectedOutput string
    }{
        {
            name:           "Testing that it works",
            key:            "password",
            stream:         bufio.NewReader(f),
            expectedOutput: "7OPecpQ8Y7CjunDbXb_FLgOrYzHtnDe0UVLLoQ1cjgc=",
        },
        {
            name:           "Testing with a different key and output succeeds",
            key:            "passworD",
            stream:         bufio.NewReader(f),
            expectedOutput: "fjLR7o07JDsDpvEGHZdA8vV2PLSxdLIL3y5uMRnk-Js=",
        },
    }

    for _, item := range cases {
        sign, _ := sign(item.key, item.stream)
        if item.expectedOutput != sign {
            t.Errorf("sign is: %v and was expecting %s", sign, item.expectedOutput)
        }
    }
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
