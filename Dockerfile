FROM golang:1.24.1-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/main.go

FROM alpine:3.22.0 

WORKDIR /app

COPY --from=build /app/app .
COPY ./configs/config.yaml .

ENV PINGROBOT_CONFIG_PATH="./config.yaml"

CMD ["./app"]


