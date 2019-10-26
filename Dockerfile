FROM golang:alpine AS build

ENV GO111MODULE=on

COPY . /go/src/github.com/dirkarnez/golang-hello-world
WORKDIR /go/src/github.com/dirkarnez/golang-hello-world

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine
COPY --from=build /go/src/github.com/dirkarnez/golang-hello-world/app /app
ENTRYPOINT [ "/usr/bin/app" ]

EXPOSE 9000