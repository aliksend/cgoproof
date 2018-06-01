FROM golang:1.9.2

ADD . /root/app
WORKDIR /root/app

RUN go run main.go
