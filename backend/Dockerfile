FROM golang:1.26.4-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

EXPOSE 8080

CMD ["app"]
