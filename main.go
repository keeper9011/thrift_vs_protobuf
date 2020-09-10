package main

import (
	"flag"
	"fmt"
	"log"
	p "thrift_vs_protobuf/protobuf/protobufmodel"
	t "thrift_vs_protobuf/thrift/thriftmodel"
)

type (
	payloadArrayType int8

	byteArray            []byte
	int64Array           []int64
	stringArray          []string
	structArrayThrift    []*t.SimpleStruct
	intStructMapThrift   map[int32]*t.SimpleStruct
	structArrayProtobuf  []*p.SimpleStruct
	intStructMapProtobuf map[int32]*p.SimpleStruct

	args struct {
		byteArraySize    int
		int64ArraySize   int
		stringArraySize  int
		structArraySize  int
		intStructMapSize int
	}
)

const (
	stringTemplate = "Lorem ipsum"

	byteArrayUsage = "[none|small|medium|huge]\n" +
		"Determines how many byte values an payload byte array will contain:\n" +
		"none    - 0 (nil value for slice)\n" +
		"small   - 4\n" +
		"medium  - 64\n" +
		"large   - 1024\n" +
		"huge    - 65536"
	int64ArrayUsage = "[none|small|medium|huge]\n" +
		"Determines how many int64 values an payload int64 array will contain:\n" +
		"none    - 0 (nil value for slice)\n" +
		"small   - 4\n" +
		"medium  - 64\n" +
		"large   - 1024\n" +
		"huge    - 65536"
	stringArrayUsage = "[none|small|medium|huge]\n" +
		"Determines how many string values (\"" + stringTemplate + "_%n%\") an payload string array will contain:\n" +
		"none    - 0 (nil value for slice)\n" +
		"small   - 4\n" +
		"medium  - 64\n" +
		"large   - 1024\n" +
		"huge    - 65536"
	structArrayUsage = "[none|small|medium|huge]\n" +
		"Determines how many struct values (int32, 4-symbol string, bool) an payload struct array will contain:\n" +
		"none    - 0 (nil value for slice)\n" +
		"small   - 4\n" +
		"medium  - 64\n" +
		"large   - 1024\n" +
		"huge    - 65536"
	intStructMapUsage = "[none|small|medium|huge]\n" +
		"Determines how many key-value pairs (int32 -> {int32, string, bool}) an payload map will contain:\n" +
		"none    - 0 (nil value for slice)\n" +
		"small   - 4\n" +
		"medium  - 64\n" +
		"large   - 1024\n" +
		"huge    - 65536"
)

const (
	smallArray  payloadArrayType = 2  // 4
	mediumArray                  = 6  // 64
	largeArray                   = 10 // 1024
	hugeArray                    = 16 // 65536

	none   = "none"
	small  = "small"
	medium = "medium"
	large  = "large"
	huge   = "huge"
)

var (
	byteArrayTypeRaw    string
	int64ArrayTypeRaw   string
	stringArrayTypeRaw  string
	structArrayTypeRaw  string
	intStructMapTypeRaw string
)

func init() {
	flag.StringVar(&byteArrayTypeRaw, "byteArray", none, byteArrayUsage)
	flag.StringVar(&int64ArrayTypeRaw, "int64Array", none, int64ArrayUsage)
	flag.StringVar(&stringArrayTypeRaw, "stringArray", none, stringArrayUsage)
	flag.StringVar(&structArrayTypeRaw, "structArray", none, structArrayUsage)
	flag.StringVar(&intStructMapTypeRaw, "intStructMap", none, intStructMapUsage)
}

func parseArraySize(raw string) int {
	switch raw {
	case none:
		return 0
	case small:
		return 1 << smallArray
	case medium:
		return 1 << mediumArray
	case large:
		return 1 << largeArray
	case huge:
		return 1 << hugeArray
	default:
		log.Fatalf("Unknown array type: %s. Values: \"%s\", \"%s\", \"%s\", \"%s\", \"%s\"",
			raw, none, small, medium, large, huge)
	}
	return 0
}

func newThriftStruct(intValue int32, stringValue string, boolValue bool) *t.SimpleStruct {
	return &t.SimpleStruct{Int32Value: &intValue, StringValue: &stringValue, BoolValue: &boolValue}
}

func newProtobufStruct(intValue int32, stringValue string, boolValue bool) *p.SimpleStruct {
	return &p.SimpleStruct{Int32Value: intValue, StringValue: stringValue, BoolValue: boolValue}
}

func makeByteArray(size, idx int) byteArray {
	if size == 0 {
		return nil
	}

	array := byteArray(make([]byte, size))
	for i := range array {
		array[i] = byte(i * idx % 256)
	}

	return array
}

func makeInt64Array(size, idx int) int64Array {
	if size == 0 {
		return nil
	}

	array := int64Array(make([]int64, size))
	for i := range array {
		array[i] = int64(i * idx)
	}

	return array
}

func makeStringArray(size, idx int) stringArray {
	if size == 0 {
		return nil
	}

	array := stringArray(make([]string, size))
	for i := range array {
		array[i] = fmt.Sprintf("%s_%d", stringTemplate, i * idx)
	}

	return array
}

func makeStructArrayThrift(size, idx int) structArrayThrift {
	if size == 0 {
		return nil
	}

	array := structArrayThrift(make([]*t.SimpleStruct, size))
	for i := range array {
		array[i] = newThriftStruct(int32(i * idx), "test", true)
	}

	return array
}

func makeIntStructMapThrift(size, idx int) intStructMapThrift {
	if size == 0 {
		return nil
	}

	table := intStructMapThrift(make(map[int32]*t.SimpleStruct, size))
	for i := 0; i < size; i++ {
		table[int32(i)] = newThriftStruct(int32(i * idx), "test", true)
	}

	return table
}

func makeStructArrayProtobuf(size, idx int) structArrayProtobuf {
	if size == 0 {
		return nil
	}

	array := structArrayProtobuf(make([]*p.SimpleStruct, size))
	for i := range array {
		array[i] = newProtobufStruct(int32(i * idx), "test", true)
	}

	return array
}

func makeIntStructMapProtobuf(size, idx int) intStructMapProtobuf {
	if size == 0 {
		return nil
	}

	table := intStructMapProtobuf(make(map[int32]*p.SimpleStruct, size))
	for i := 0; i < size; i++ {
		table[int32(i)] = newProtobufStruct(int32(i * idx), "test", true)
	}

	return table
}

func createThriftPayload(args args, idx int) *t.Payload {
	idx++
	return &t.Payload{
		ByteArray:    makeByteArray(args.byteArraySize, idx),
		Int64Array:   makeInt64Array(args.int64ArraySize, idx),
		StringArray:  makeStringArray(args.stringArraySize, idx),
		StructArray:  makeStructArrayThrift(args.stringArraySize, idx),
		IntStructMap: makeIntStructMapThrift(args.intStructMapSize, idx),
	}
}

func createProtobufPayload(args args, idx int) *p.Payload {
	idx++
	return &p.Payload{
		ByteArray:    makeByteArray(args.byteArraySize, idx),
		Int64Array:   makeInt64Array(args.int64ArraySize, idx),
		StringArray:  makeStringArray(args.stringArraySize, idx),
		StructArray:  makeStructArrayProtobuf(args.stringArraySize, idx),
		IntStructMap: makeIntStructMapProtobuf(args.intStructMapSize, idx),
	}
}

func main() {
	flag.Parse()

	args := args{
		parseArraySize(byteArrayTypeRaw),
		parseArraySize(int64ArrayTypeRaw),
		parseArraySize(stringArrayTypeRaw),
		parseArraySize(structArrayTypeRaw),
		parseArraySize(intStructMapTypeRaw),
	}

	log.Printf("Using payload: [%d]byte, [%d]int64, [%d]string, [%d]struct, len(map[int32]struct)=%d",
		args.byteArraySize,
		args.int64ArraySize,
		args.stringArraySize,
		args.stringArraySize,
		args.intStructMapSize)

	msgThrift := createThriftPayload(args, 0)
	msgProtobuf := createProtobufPayload(args, 0)

	log.Println("Messages created")

	payloadThrift := serializeThrift(msgThrift)
	payloadProtobuf := serializeProtobuf(msgProtobuf)

	log.Printf("Messages serialized.\n> thrift size %d bytes\n> protobuf size %d bytes",
		len(*payloadThrift), len(*payloadProtobuf))

	newMsgThrift := t.Payload{}
	deserializeThrift(payloadThrift, &newMsgThrift)
	newMsgProtobuf := p.Payload{}
	deserializeProtobuf(payloadProtobuf, &newMsgProtobuf)

	log.Println("Messages deserialized")
}
