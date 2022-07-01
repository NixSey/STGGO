# syntax=docker/dockerfile:1
FROM debian:10 as kontol
RUN ls
RUN apt-get install go -y
RUN printf "500" | go run main.go
