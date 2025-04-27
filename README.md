# Go CRUD API with Docker and Makefile

This project is a simple **Go CRUD API** that uses **Gin** and **PostgreSQL**. It is configured with Docker for containerized development, and a `Makefile` is provided for simplifying common tasks like building, testing, and running the application.

## Table of Contents
- [Technologies](#technologies)
- [Setup](#setup)
- [Development](#development)
- [Running with Docker](#running-with-docker)
- [Makefile Commands](#makefile-commands)
- [Testing](#testing)
- [Cleanup](#cleanup)

## Technologies
- **Go (1.24.2)**: The programming language used for the API.
- **Gin**: Web framework used to build the RESTful API.
- **PostgreSQL**: Database used for storing data.
- **Docker**: Containerization for the app and database.
- **Makefile**: For automating common tasks.
- **CompileDaemon**: For automatic rebuilding and reloading the Go server during development.

## Setup

### Prerequisites
1. Install **Go (1.24.2)** on your machine. You can download Go from [here](https://golang.org/dl/).
2. Install **Docker**. Follow the instructions on the official [Docker website](https://docs.docker.com/get-docker/).
3. Install **Make** (optional but recommended for running commands from the `Makefile`).

### Install Dependencies

1. Clone the repository:

bash 

```git clone <repository-url>```

```cd go-crud  ```

2.	Install Go dependencies:

```go mod tidy```

Install CompileDaemon (for local development)

To install CompileDaemon globally on your system:

```go install github.com/githubnemo/CompileDaemon@latest```

Ensure $GOPATH/bin is in your PATH. To make it permanent, add the following line to your shell profile (~/.bashrc or ~/.zshrc):

```export PATH=$PATH:$(go env GOPATH)/bin```

Then run:

```source ~/.bashrc  # or ~/.zshrc```

Docker Setup

1.	You will also need Docker Compose to run the database container. Ensure Docker and Docker Compose are installed on your machine.
2.	Set up the database with Docker by using docker-compose:

docker-compose -f docker-compose.db.yml up


This will start a PostgreSQL database running in a container.

Development

Running the Server Locally

1.	Start the database container with:

```make docker-compose-db```

2.	Start the Go server with live reloading using CompileDaemon:

```make dev```

This command will:
	•	Watch for changes in your Go code and automatically rebuild the app.
	•	Restart the Go server when changes are detected.

Using Makefile Commands

Here are the available commands in the Makefile:

build

Build the Go server binary:

```make build```

This will create a go-crud binary in the current directory.

run

Run the Go server without building:

```make run```

This will directly run the main.go file.

dev

Run the server with automatic reloading during development:

```make dev```

This will use CompileDaemon to monitor for changes in the Go files and automatically rebuild and restart the server.

test

Run the tests:

```make test```

lean

Clean up build artifacts:

```make clean```

docker-compose

Start the entire application (including the database) using Docker Compose:

```make docker-compose```

This will bring up both the database and the Go app, rebuilding as needed.

docker-compose-db

Start only the database using Docker Compose:

```make docker-compose-db```


Notes
	•	The project is set up to use PostgreSQL as the database, and the application depends on the database being up and running for most operations.
	•	During development, CompileDaemon is used to automatically rebuild and reload the Go server when files change.
	•	You can customize the Docker Compose configuration if you need to modify the database connection or other services.