FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/main.go

FROM alpine:latest

WORKDIR /root
COPY --from=build /app/main .
COPY --from=build /app/.env .

EXPOSE 3000

CMD ["./main"]
