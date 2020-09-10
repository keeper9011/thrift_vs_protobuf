namespace go thriftmodel

struct SimpleStruct
{
	1:	optional	i32						int32Value,
	2:	optional	string					stringValue,
	3:	optional	bool					boolValue,
}

struct Payload
{
	1:	optional	binary					byteArray,
	2:	optional	list<i64>				int64Array,
	3:	optional	list<string>			stringArray,
	4:	optional	list<SimpleStruct>		structArray,
	5:	optional	map<i32, SimpleStruct>	intStructMap,

}