package main

import (
	"log"
	"mmrp-scraper/internal/scrapers"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// TODO:
// 1. Парсить таблицу из MAPM, прикрутить авторизацю на сайте
// 2. Создать эндпоинт, через который можно пинговать сервер и проверять работу (сделать так, чтобы при пинге выдавались логи)
// 3. Настроить graceful shutdown
// 4. Через профилировщик проверить на ненужные аллокации, оптимизация кода
// 5. Оптимизация докер-контейнера
// 6. Проверить код через линтеры
// 7. Если сервер упал (сайт поменял домен, ошибки в веб-парсинге и т.д.), то сервис не должен перезапускаться
// 8. Настроить CI/CD
// 9. Написать тесты

func run() error {
	// Инициализация парсеров
	sMmrp := scrapers.MMRPScraper{}
	sMapm := scrapers.MAPMScraper{}
	cs, err := scrapers.Init()
	if err != nil {
		log.Fatal(err, "Error in configs for scrapper")
	}
	// Создание кронов
	scheduler := gocron.NewScheduler(time.UTC)

	// Скраппинг MMRP каждые 15 минут
	scheduler.Every(15).Minutes().Do(func() {
		log.Println("Start MMRP task")
		sMmrp.Scrape(cs)
	})

	// Скраппинг MAPM каждые 10 часов
	scheduler.Every(10).Hours().Do(func() {
		log.Println("Start MAPM task")
		sMapm.ScrapeWithRod()
	})
	scheduler.StartBlocking()
	return nil
}
