BINARY=smartgrid.exe

all:	dependencies
	@echo "Build project :"
	@mkdir -p bin
	go build -o ./bin/$(BINARY) ./src/smartgrid.go
	@echo "SUCCESFULL COMPILATION"
	@echo "\n\tRun ./bin/$(BINARY)\n"
	

clean:
	@echo "Removing binaries..."
	@rm -rf ./bin

re : clean all

dependencies:
	@echo "Downloading dependencies..."
	#go get -u github.com/fatih/color

test: re
	@echo "Auto run $(BINARY)."
	@./bin/$(BINARY)