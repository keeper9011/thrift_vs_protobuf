package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func serializeProtobuf(msg proto.Message) *[]byte {
	payload, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("Protobuf serialize failed: %v", err)
	}
	return &payload
}

func deserializeProtobuf(payload *[]byte, msg proto.Message) {
	err := proto.Unmarshal(*payload, msg)
	if err != nil {
		log.Fatalf("Protobuf deserialize failed: %v", err)
	}
}