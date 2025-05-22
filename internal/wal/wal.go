package wal

import (
	"fmt"
	"log"

	protoapi "github.com/shohagrana006/wal/pkg/api/wal"
	"google.golang.org/protobuf/proto"
)

func Wal(){
	req := &protoapi.HelloRequest{
		Name: "Rana",
	}
	result := SayHello(req)
	data, err := proto.Marshal(result)
	if err != nil{
		log.Fatalf("Marshing err %v", err)
	}
	fmt.Print("Marshing - ",data)

	a := &protoapi.HelloResponse{}
	err = proto.Unmarshal(data, a)
	if err != nil {
		log.Fatalf("Unmarshing err %v", err)
	}

	fmt.Printf("Unmarshing data - %v", a)

}


func SayHello(req *protoapi.HelloRequest) *protoapi.HelloResponse{
	return &protoapi.HelloResponse{
		Message: "Hello Shohag " + req.Name + " !!",
	}
}