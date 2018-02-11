test:
	go test ./...
	
debug:
	go build

release:
	go build -gcflags=-B
