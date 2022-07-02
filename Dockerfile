# syntax=docker/dockerfile:1
FROM debian:10 as kontol
RUN apt update && apt upgrade -y &> /dev/null
RUN apt-get install curl wget -y &> /dev/null
RUN wget --quiet "https://dl.google.com/go/$(curl -L 'https://golang.org/VERSION?m=text').linux-amd64.tar.gz"
RUN tar xvf go*.linux-amd64.tar.gz
RUN chown -R root:root ./go
RUN ls
RUN mv go /usr/local
RUN echo "export GOROOT=\$HOME/go\nexport GOPATH=\$HOME/work\nexport PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> ~/.profile
RUN printf "500" | go run main.go
