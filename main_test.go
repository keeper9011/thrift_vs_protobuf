package main

import (
	"testing"
	"thrift_vs_protobuf/protobuf/protobufmodel"
	"thrift_vs_protobuf/thrift/thriftmodel"
)

func BenchmarkSerializeThrift(b *testing.B) {
	args := []struct {
		name string
		args
	}{
		{
			none,
			args{
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
			},
		},
		{
			small,
			args{
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
			},
		},
		{
			medium,
			args{
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
			},
		},
		{
			large,
			args{
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
			},
		},
		{
			huge,
			args{
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
			},
		},
	}

	for _, arg := range args {
		b.Run(arg.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				msg := createThriftPayload(arg.args, i)
				b.StartTimer()

				serializeThrift(msg)
			}
		})
	}
}

func BenchmarkDeserializeThrift(b *testing.B) {
	args := []struct {
		name string
		args
	}{
		{
			none,
			args{
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
			},
		},
		{
			small,
			args{
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
			},
		},
		{
			medium,
			args{
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
			},
		},
		{
			large,
			args{
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
			},
		},
		{
			huge,
			args{
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
			},
		},
	}

	for _, arg := range args {
		b.Run(arg.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				msg := createThriftPayload(arg.args, i)
				payload := serializeThrift(msg)
				newMsg := thriftmodel.Payload{}
				b.StartTimer()

				deserializeThrift(payload, &newMsg)
			}
		})
	}
}

func BenchmarkSerializeProtobuf(b *testing.B) {
	args := []struct {
		name string
		args
	}{
		{
			none,
			args{
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
			},
		},
		{
			small,
			args{
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
			},
		},
		{
			medium,
			args{
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
			},
		},
		{
			large,
			args{
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
			},
		},
		{
			huge,
			args{
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
			},
		},
	}

	for _, arg := range args {
		b.Run(arg.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				msg := createProtobufPayload(arg.args, i)
				b.StartTimer()

				serializeProtobuf(msg)
			}
		})
	}
}

func BenchmarkDeserializeProtobuf(b *testing.B) {
	args := []struct {
		name string
		args
	}{
		{
			none,
			args{
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
				parseArraySize(none),
			},
		},
		{
			small,
			args{
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
				parseArraySize(small),
			},
		},
		{
			medium,
			args{
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
				parseArraySize(medium),
			},
		},
		{
			large,
			args{
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
				parseArraySize(large),
			},
		},
		{
			huge,
			args{
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
				parseArraySize(huge),
			},
		},
	}

	for _, arg := range args {
		b.Run(arg.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				msg := createProtobufPayload(arg.args, i)
				payload := serializeProtobuf(msg)
				newMsg := protobufmodel.Payload{}
				b.StartTimer()

				deserializeProtobuf(payload, &newMsg)
			}
		})
	}
}