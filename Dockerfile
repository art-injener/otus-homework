FROM golang:1.16-alpine3.14 AS builder

RUN go version

COPY . /github.com/art-injener/otus
WORKDIR /github.com/art-injener/otus

RUN go mod download

RUN GOOS=linux  go build -o ./bin/server ./cmd/webserver/main.go

FROM alpine:latest
LABEL maintainer="Artem Danilchenko <art_injener@mail.ru>"

WORKDIR /root/

COPY --from=0 /github.com/art-injener/otus/bin/server .
COPY --from=0 /github.com/art-injener/otus/configs configs/

RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime

ENTRYPOINT ["./server"]
