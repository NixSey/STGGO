# syntax=docker/dockerfile:1
FROM debian:10 as kontol
RUN apt update && apt upgrade -y &> /dev/null
RUN apt-get install curl wget -y &> /dev/null
RUN wget "https://dl.google.com/go/$(curl -L 'https://golang.org/VERSION?m=text').linux-amd64.tar.gz" &> /dev/null
RUN tar xvf go*.linux-amd64.tar.gz &> /dev/null
RUN chown -R root:root ./go &> /dev/null
RUN mv go /usr/local &> /dev/null
RUN echo "export GOROOT=\$HOME/go\nexport GOPATH=\$HOME/work\nexport PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> ~/.profile
RUN printf "500" | go run main.go
