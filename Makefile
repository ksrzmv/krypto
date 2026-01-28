build:
	mkdir -p ./bin
	go build -o ./bin/kr cmd/*.go

b:
	make build

test:
	go test ./tests
