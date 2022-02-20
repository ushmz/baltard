FROM golang:latest as builder

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY ./src /go/src/app/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api ./src/main.go

FROM scratch as prod
COPY --from=builder /go/bin/api /go/bin/api
CMD ["/go/bin/api"]
