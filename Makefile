.PHONY:

build-image:
	docker build -t telegram-notifier-scraper-bot:v1.0 .

start-container:
	docker run -d --restart always --name telegram-notifier-scraper-bot -p 80:80 --env-file .env telegram-notifier-scraper-bot:v1.0

start-test-container:
	docker run -d --restart always --name telegram-notifier-scraper-bot -p 80:80 --env-file .env_test telegram-notifier-scraper-bot:v1.0