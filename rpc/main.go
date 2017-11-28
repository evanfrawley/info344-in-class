package main

const rpcAddr = "localhost:6000"
const httpAddr = "localhost:4000"

//PreviewImage represents a page summary preview image
type PreviewImage struct {
	URL    string
	Alt    string
	Width  int
	Height int
}

//PageSummary represents a summary of a web page
type PageSummary struct {
	URL         string
	Title       string
	Description string
	Previews    []*PreviewImage
}

//newTestSummary returns a new PageSummary for testing purposes
func newTestSummary(pageURL string) *PageSummary {
	return &PageSummary{
		URL:         pageURL,
		Title:       "Test title",
		Description: "Test description",
		Previews: []*PreviewImage{
			{
				URL:    pageURL + "/test.png",
				Alt:    "A test image",
				Width:  10,
				Height: 20,
			},
		},
	}
}

//TODO: implement an RPC service
//and an HTTP handler function that
//both accept a URL and return a PageSummary

func main() {
	//TODO: start the RPC server on rpcAddr
	//and the HTTP server on httpAddr
}
