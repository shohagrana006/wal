package main

import (
	"testing"

	protoapi "github.com/shohagrana006/wal/pkg/api/wal"
)

func TestSayHello(t *testing.T) {
	req := &protoapi.HelloRequest{Name: "Shohag"}
	resp := SayHello(req)
	expected := "Hello, Shohag!"

	if resp.Message != expected {
		t.Errorf("Expected %s, got %s", expected, resp.Message)
	}
}
