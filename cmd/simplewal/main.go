package main

import (
	"os"

	"github.com/shohagrana006/simplewal/internal/wal"
)

func main(){
	if len(os.Args) > 1 && os.Args[1] == "read"{
		wal.Read()
	}else{
		wal.Write();
	}
}