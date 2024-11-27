FROM golang:1.23.3-bookworm AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -o /bin/mqtt-gateway

FROM --platform=linux/arm64/v8 arm64v8/alpine:latest

COPY --from=build /bin/mqtt-gateway /bin

WORKDIR /app
EXPOSE 1883
CMD ["/bin/mqtt-gateway"]
