package main

import (
	"testing"
	"github.com/shohagrana006/simplewal/pkg/api"
)

func TestSayHello(t *testing.T) {
	req := &protoapi.HelloRequest{Name: "Shohag"}
	resp := SayHello(req)
	expected := "Hello, Shohag!"

	if resp.Message != expected {
		t.Errorf("Expected %s, got %s", expected, resp.Message)
	}
}
