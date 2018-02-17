GOFLAGS = -gcflags -B
PACKAGE = github.com/dbriemann/chesskimo/cmd

test:
	go test ./...
	
clean:
	rm chesskimo bench
	
debug:
	go build -o chesskimo -i -v -ldflags="-X main.version=$(shell git describe --always)" $(PACKAGE)/chesskimo
	go build -o bench -i -v -ldflags="-X main.version=$(shell git describe --always)" $(PACKAGE)/bench

release:
	go build -o chesskimo -i -v -gcflags="-B" -ldflags="-X main.version=$(shell git describe --always)" $(PACKAGE)/chesskimo
	go build -o bench -i -v -gcflags="-B" -ldflags="-X main.version=$(shell git describe --always)" $(PACKAGE)/bench

profile:
	./bench -profile=prof.out
	go tool pprof -callgrind -output=profile.grind bench prof.out