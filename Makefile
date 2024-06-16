build:
	@go build -o ./bin/imc-server ./main.go
run: build
	@./bin/imc-server
test:
	@go test -v ./...