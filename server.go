package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/new-endpoint", handleNewWndpoint)

	fmt.Println("Add new prirnt...")


	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil{
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request){
	if request.Method!= "GET"{
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writerResponse(writer, "Hello World!")
	
}

func handleHealth(writer http.ResponseWriter, request *http.Request){
	if request.Method!= "GET"{
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writerResponse(writer, "OK!")
}
func handleNewWndpoint(writer http.ResponseWriter, request *http.Request){
	if request.Method!= "GET"{
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writerResponse(writer, "OK!")
	
}
func writerResponse(writer http.ResponseWriter, responseString string){
	response := []byte(responseString)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
