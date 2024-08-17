BINARY_NAME=todo

build:
	go build -o ./bin/${BINARY_NAME} cmd/todo/main.go

run: 
	./bin/${BINARY_NAME} ${ARGS}