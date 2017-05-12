pb: server.proto
	mkdir pb && protoc --go_out=plugins=grpc:pb server.proto
