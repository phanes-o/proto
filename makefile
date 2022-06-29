
build:
	protoc -I=. --go_out=. --micro_out=. primitive/primitive.proto
	protoc -I=. --go_out=. --micro_out=. example/user.proto