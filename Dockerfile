FROM golang:1.24-alpine

WORKDIR /app

# Copy and download dependencies first (leveraging Docker layer cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code and build
COPY . .
RUN go build -o account-service main.go

EXPOSE 8080
CMD ["./account-service"]
