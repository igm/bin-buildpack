package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	go func() {
		cmd := exec.Command("./server")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		cmd.Wait()
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port: ", port)
	u, err := url.Parse("http://localhost")
	if err != nil {
		log.Fatal(err)
	}
	rp := httputil.NewSingleHostReverseProxy(u)

	tr := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.Dial("unix", "/tmp/server.sock")
		},
	}
	rp.Transport = tr
	http.Handle("/", rp)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handlerHW(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World!"))
}
