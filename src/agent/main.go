package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port: ", port)
	http.HandleFunc("/", handlerHW)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func handlerHW(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World!"))
}
