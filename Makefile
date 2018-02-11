test:
	go test ./...
	
debug:
	go build -o chesskimo -i -v -ldflags="-X main.version=$(git describe --always --long --dirty)" github.com/dbriemann/chesskimo/cmd/chesskimo
	go build -o bench -i -v -ldflags="-X main.version=$(git describe --always --long --dirty)" github.com/dbriemann/chesskimo/cmd/bench

release:
	go build -o chesskimo -i -v -gcflags="-B" -ldflags="-X main.version=$(git describe --always --long --dirty)" github.com/dbriemann/chesskimo/cmd/chesskimo
	go build -o bench -i -v -gcflags="-B" -ldflags="-X main.version=$(git describe --always --long --dirty)" github.com/dbriemann/chesskimo/cmd/bench
