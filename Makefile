compile:
	protoc api/v1.proto --go_out=. --go-opt=paths=source_relative --proto-path=.

test:
	go test -race ./..