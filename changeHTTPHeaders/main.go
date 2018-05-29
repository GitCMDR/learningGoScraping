package main

import (
	"net/http"
	"time"
	"log"
	"os"
	"io"
	"fmt"
)

// One of the reasons I may decide to change Headers
// is to identify my web scraper and manage to have
// good relationships with the webmasters of the sites
// i'm scraping from. It allows for the bot to easily
// be throttled or blocked in case the bot causes issues.

func main() {
	// Create a HTTP client with Timeout
	client := &http.Client {
		Timeout: 30 * time.Second,
	}

	// Define a Request in order to change the Header but
	// not send it yet.
	request, err := http.NewRequest("GET", "http://webscraper.io/test-sites/e-commerce/allinone", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// Now that the Request is defined, override the Header!
	request.Header.Set("User-Agent", "MyScraperBot v1.0 https://www.github.com/gitCMDR/")

	fmt.Println(request)
	// Now let's make a Request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()


	// Print the response to StdOut
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

}


