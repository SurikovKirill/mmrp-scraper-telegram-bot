FROM golang:1.17-alpine3.13 AS builder

RUN go version

COPY . /mmrp-scraper/
WORKDIR /mmrp-scraper/

RUN go mod tidy
RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/mmrp/main.go

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
RUN apk add --no-cache chromium
RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
WORKDIR /root/

COPY --from=0 /mmrp-scraper/.bin/bot .

EXPOSE 80

CMD ["./bot"]
