package main

import (
	"fmt"
//	"log"
	"net/http"
	
//	"github.com/gorilla/mux"
	
	
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!!\n"))
}

func main() {
	fmt.Println("Start Restful api for file information!")
	
	
//	r := mux.NewRouter()
//
//	r.HandleFunc("/", yourHandler)
//	
//	// Bind to a port and pass our router in
//    log.Fatal(http.ListenAndServe(":3030", r))
}

