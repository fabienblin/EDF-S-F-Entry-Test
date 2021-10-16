EMS=ems.exe
ESS=ess.exe
POC=poc.exe
PV=pv.exe
SMARTGRID=smartgrid.exe

all:	dependencies
	mkdir -p bin
	go build -o ./bin/$(EMS) ./src/ems/ems.go
	# go build -o ./bin/$(ESS) ./src/ess/ess.go
	go build -o ./bin/$(POC) ./src/poc/poc.go
	go build -o ./bin/$(SMARTGRID) ./src/smartgrid/smartgrid.go
	#go build -o ./bin/$(PV) ./src/pv/pv.go
	@echo "SUCCESFULL COMPILATION"
	@echo "run ./bin/ems.exe"
	

clean:
	rm -rf ./bin

re : clean all

dependencies:
#	go get -u gonum.org/v1/gonum
#	go get -u gonum.org/v1/plot

test: re
	./bin/$(EMS)