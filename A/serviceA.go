package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Service A running Port 8082")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8082", nil)
}

func healthz(r http.ResponseWriter, w *http.Request) {
	// Just return HTTP 200 to indicate it's healthy
	r.WriteHeader(http.StatusOK)
}

func respond(r http.ResponseWriter, w *http.Request) {
	service, ok := os.LookupEnv("MICROSERVICE")
	if !ok {
		service = "localhost:8081"
	}

	var result string = "#ERROR"
	response, err := http.Get(fmt.Sprintf("http://%v", service))
	
	
	if err == nil {
		body, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err == nil && response.StatusCode == 200 {
			result = string(body)	
		} else {
			fmt.Printf("### Unable to read Body: %v\n", err)
		}
	} else {
		fmt.Printf("### Unable to Reach: http://%v\n", service)
	}

	version, ok := os.LookupEnv("VERSION")
	if !ok {
		version = ""
	}

	r.Write([]byte(fmt.Sprintf("A(%v) --> %v", version, result)))
}