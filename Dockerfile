# syntax=docker/dockerfile:1
FROM golang:1.16 as kontol
RUN ls
RUN printf "500" | go run main.go
