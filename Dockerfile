# syntax=docker/dockerfile:1
FROM debian:latest
RUN apt-get install go -y
CMD printf "500" | go run main.go
