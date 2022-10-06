.PHONY:

build:
	go build -o ./.bin/bot cmd/mmrp/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-notifier-scraper-bot:v0.1 .

start-container:
	docker run -d --restart always --name telegram-notifier-scraper-bot -p 80:80 --env-file .env telegram-notifier-scraper-bot:v0.1

start-container-test:
	docker run -d --restart always --name telegram-notifier-scraper-bot -p 80:80 --env-file .env_test telegram-notifier-scraper-bot:v0.1