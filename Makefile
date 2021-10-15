BINARY=EMS.exe

env:
	@export GOPATH=${pwd}

all:	env dependencies
	mkdir -p bin
	go build -o ./bin/$(BINARY) ./src/main.go

clean:
	rm -rf ./bin

re : clean all

dependencies:
#	go get -u gonum.org/v1/gonum
#	go get -u gonum.org/v1/plot

test: re
	./bin/$(BINARY)