build:
	go build -o myapp main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf myapp