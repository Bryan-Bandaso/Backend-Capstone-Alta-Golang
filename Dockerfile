FROM golang:1.17.8

WORKDIR /go/Backend-Capstone-Alta-Golang
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080
CMD ["go", "run", "main.go"]