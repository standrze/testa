package main

import (
	"darah/auto"
	"fmt"
	"net/http"
)

func main() {
	run()
}

func run() {
	first := &auto.Sling{
		HttpMethod: http.MethodGet,
		Path: "/",
		Port: 80,
		SSL: false,
		Header: make(http.Header),
	}
	second := &auto.Sling{
		HttpMethod: http.MethodPost,
		Path: "/",
		Port: 80,
		SSL: false,
		Header: make(http.Header),
	}

	request, err := first.Request("127.0.0.1")
	if err != nil {
		fmt.Println("Fug error")
	}

	fmt.Println(request)

	request, err = second.Request("127.0.0.1")
	if err != nil {
		fmt.Println("Fug error")
	}

	fmt.Println(request)
}