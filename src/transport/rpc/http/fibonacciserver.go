package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"shared/parameters"
	"apps/fibonacci/implrpc"
	"shared/shared"
)

func main() {

	// create new instance of fibonacci
	fibonacci := new(implrpc.Fibonacci)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Fibonacci", fibonacci)

	// associate a http handler to server
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, e := net.Listen("tcp", shared.ResolveHostIp()+":"+strconv.Itoa(parameters.FIBONACCI_PORT))
	if e != nil {
		log.Fatal("Server not started:", e)
	}

	// wait for calls
	fmt.Println("Server is running... \n")
	http.Serve(l, nil)
}
