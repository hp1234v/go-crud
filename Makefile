build:
	go build -o server main.go

run:
	go run main.go

install-compileDaemon:
	# Install CompileDaemon
	go install github.com/githubnemo/CompileDaemon@latest

set-path:
	# Ensure GOPATH/bin is added to PATH
	export PATH=$(go env GOPATH)/bin:$PATH

check-daemon:
	# Check if CompileDaemon is in the path
	which CompileDaemon || echo "CompileDaemon not installed correctly."

dev: 
	# Run CompileDaemon
	CompileDaemon --build="go build -o server main.go" --command="./server"

test:
	go test ./...

clean:
	rm -rf server

docker-compose:
	docker compose down -v && docker compose up --build

docker-compose-db:
	docker compose -f docker-compose.db.yml down -v && docker compose -f docker-compose.db.yml up