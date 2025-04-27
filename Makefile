build:
	go build -o myapp main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf myapp

docker-compose:
	docker compose down -v && docker compose up --build