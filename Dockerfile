# syntax=docker/dockerfile:1
FROM golang:1.16-alpine as kontol
WORKDIR /ppk
COPY go.mod ./
copy go.sum ./
RUN go mod download
COPY *.go ./
RUN go build main-cli.go
EXPOSE 8080
CMD ["/main-cli 500"]
