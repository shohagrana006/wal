package wal

import (
	"bufio"
	"fmt"
	"log"
	"os"

	userproto "github.com/shohagrana006/wal/pkg/api/user"
	"google.golang.org/protobuf/proto"
)

func Write(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write your name: ")
	name, err := reader.ReadString('\n')
	if nil != err {
		log.Fatalf("Read err %v ", err)
	}
	name = name[:len(name)-1]

	fmt.Print("Write your email: ")
	email, err := reader.ReadString('\n')
	if nil != err {
		log.Fatalf("Read err %v ", err)
	}
	email = email[:len(email)-1]

	req := &userproto.UserRequest{
		Name: name,
		Email: email,
	}
	data, err := proto.Marshal(req)
	if nil != err{
		log.Fatalf("Marshal error: %v", err)
	}
	// fmt.Print(data)

	er := os.WriteFile("internal/storage/storage.data", data, 0644)
	if nil != er {
		log.Fatalf("Failed to write file: %v", er)
	}
	fmt.Println("Write successfully to storage.data")
}

func Read(){
	data, err := os.ReadFile("internal/storage/storage.data")
	if nil != err{
		log.Fatalf("Data cannot read, err is %v", err)
	}
	res := &userproto.UserResponse{}
	
	err = proto.Unmarshal(data,res)
	if nil != err{
		log.Fatalf("Data unmarshing err is %v", err)
	}
	fmt.Printf("Name: %v \nEmail: %v\n", res.Name, res.Email)
}
