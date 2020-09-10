package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
)

const BufferSize = 1024 * 128

func serializeThrift(msg thrift.TStruct) *[]byte {
	t := thrift.NewTMemoryBufferLen(BufferSize)
	p := thrift.NewTCompactProtocolFactory().GetProtocol(t)
	serializer := &thrift.TSerializer{Transport: t, Protocol: p}

	payload, err := serializer.Write(nil, msg)
	if err != nil {
		log.Fatalf("Thrift serialize failed: %v", err)
	}
	return &payload
}

func deserializeThrift(payload *[]byte, msg thrift.TStruct) {
	t := thrift.NewTMemoryBufferLen(BufferSize)
	p := thrift.NewTCompactProtocolFactory().GetProtocol(t)
	deserializer := &thrift.TDeserializer{Transport: t, Protocol: p}

	_ = deserializer.Transport.Close() // resets underlying bytes.Buffer

	err := deserializer.Read(msg, *payload)
	if err != nil {
		log.Fatalf("Thrift deserialize failed: %v", err)
	}
}
