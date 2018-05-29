package main

import (
	"net/http"
	"log"
	"os"
	"io"
)

func main() {
	// Make HTTP Request to Test Site (http://webscraper.io/test-sites/e-commerce/allinone)
	response, err := http.Get("http://webscraper.io/test-sites/e-commerce/allinone")
	// Check for errors.
	if err != nil {
		log.Fatal(err)
	}

	// Close the Body of the response just before terminating the program.
	defer response.Body.Close()

	// Pass the response to Standard Output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print number of bytes on string - Realistically you can throw this away by saving it in '_' variable.
	log.Println("Amount of bytes in StOut:", n)
}