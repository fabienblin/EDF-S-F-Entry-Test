all:	dependencies
	mkdir -p bin
	go build -o bin/EMS.exe src/main.go

clean:
	rm -rf ./bin

re : clean all

dependencies:
#	go get -u gonum.org/v1/gonum
#	go get -u gonum.org/v1/plot