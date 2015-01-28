package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerHello)
	l, err := net.Listen("unix", "/tmp/server.sock")
	defer func() {
		l.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(l, http.DefaultServeMux)
	log.Fatal(err)
}

func handlerHello(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World!"))
}
