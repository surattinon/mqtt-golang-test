FROM golang:1.23.3-bookworm AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# MQTT Sub Test
RUN go run main.go & \
	sleep 5 && \
	go test -v ./... && \
	pkill -f main.go

RUN GOOS=linux GOARCH=arm64 go build -o /bin/mqtt-gateway

FROM --platform=linux/arm64/v8 arm64v8/alpine:latest

COPY --from=build /bin/mqtt-gateway /bin

WORKDIR /app
EXPOSE 1883 8080
CMD ["/bin/mqtt-gateway"]
