package wal

import "simplewal/proto"

func AppendEntry(msg string) *proto.LogEntry {
	return &proto.LogEntry{Message:msg}
}