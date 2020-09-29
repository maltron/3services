package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Service B running Port 8081")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8081", nil)
}

func healthz(r http.ResponseWriter, w *http.Request) {
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

	r.Write([]byte(fmt.Sprintf("B --> %v", result)))
}