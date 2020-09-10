go test -bench=^Benchmark.*Thrift$ ./... > thrift.txt
go test -bench=^Benchmark.*Protobuf$ ./... > protobuf.txt
