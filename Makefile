pb:
	protoc -I=. --gogofaster_out=:./ --gogofaster_opt=paths=source_relative ./*.proto