FROM golang:1.16.7

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

EXPOSE 8080
