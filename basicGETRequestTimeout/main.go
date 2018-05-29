package main

import (
	"net/http"
	"time"
	"log"
	"os"
	"io"
)

// It is possible to set up a 'Timeout' parameter within your HTTP
// Client. You just need to define a HTTP Client (as below!)

// remember http.Get uses the default HTTP client with default settings
// and that's why we need to override it!

func main() {
	// Creating HTTP Client
	client := &http.Client {
		Timeout: 30 * time.Second,  // define a timeout of 30 seconds (30,000 ms)
	}

	// Now lets make the request!
	response, err := client.Get("http://webscraper.io/test-sites/e-commerce/allinone")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()  // close the Body just before finishing program execution

	// Just as before - let's send the response to StdOut
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to StdOut:", n)

}