package main

import (
	"fmt"

	protoapi "github.com/shohagrana006/wal/pkg/api/wal"
)

func main() {
	req := &protoapi.HelloRequest{Name: "Shohag"}
	resp := SayHello(req)
	fmt.Println(resp.Message)
}

func SayHello(req *protoapi.HelloRequest) *protoapi.HelloResponse {
	return &protoapi.HelloResponse{
		Message: "Hello, " + req.Name + "!",
	}
}
