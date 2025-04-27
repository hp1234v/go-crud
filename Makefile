build:
	go build -o go-crud main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf myapp

docker-compose:
	docker compose down -v && docker compose up --build

docker-compose-db:
	docker compose -f docker-compose.db.yml down -v && docker compose -f docker-compose.db.yml up