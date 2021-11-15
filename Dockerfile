FROM golang:1.17-alpine3.13 AS builder

RUN go version

COPY . /mmrp-scraper/
WORKDIR /mmrp-scraper/

RUN go mod tidy
RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/mmrp/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /mmrp-scraper/.bin/bot .
COPY --from=0 /mmrp-scraper/configs configs/

EXPOSE 80

CMD ["./bot"]
