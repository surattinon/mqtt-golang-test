# Simple MQTT to API with Go

## Dependencies

- Go (Golang)
- Docker (optional, for containerized environment)

### Clone the repository:

```bash
git clone https://github.com/surattinon/mqtt-golang-test
cd mqtt-golang-test
```

### Setup and Build (Without Docker)

1. Install dependencies and tidy up the Go module:

   ```bash
   go mod tidy
   ```

1. Build the project and run the application:
   ```bash
   go build -o bin/mqtt && ./bin/mqtt
   ```

### Setup and Build (With Docker)

1. If you prefer using Docker for the build and deployment process, you can use Docker Compose:
   ```bash
   docker compose up -d
   ```
