FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/api
RUN go build -o /app/main .

CMD ["/app/main"]