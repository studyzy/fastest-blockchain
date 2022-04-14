pb:
	protoc -I=. --gogofaster_out=:./ --gogofaster_opt=paths=source_relative ./*.proto

rpc:
	protoc -I=. \
		--gogofaster_out=plugins=grpc:./ \
		--gogofaster_opt=paths=source_relative \
		*.proto