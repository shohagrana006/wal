package main

import (
	"fmt"
	"github.com/shohagrana006/simplewal/pkg/api"
)

func main() {
	req := &protoapi.HelloRequest{Name: "Shohag"}
	resp := SayHello(req)
	fmt.Println(resp.Message)
}

func SayHello(req *protoapi.HelloRequest) *protoapi.HelloReply {
	return &protoapi.HelloReply{
		Message: "Hello, " + req.Name + "!",
	}
}
