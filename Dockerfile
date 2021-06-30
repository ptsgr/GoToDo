FROM golang:1.15-buster

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o todo ./cmd/main.go

CMD ["./todo"]