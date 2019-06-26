all: cli1 cli2

cli1: cli1/main.go
	go build -o bin/cli1 cli1/main.go

cli2: cli2/main.go
	go build -o bin/cli2 cli2/main.go

clean:
	rm -rf bin
