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
// 4. Создать эндпоинт, через который можно пинговать сервер и проверять работу (сделать так, чтобы при пинге выдавались логи)
// 5. Настроить graceful shutdown
// 6. Сделать логгирование в файл
// 7. Через профилировщик проверить на ненужные аллокации, оптимизация кода
// 8. Оптимизация докер-контейнера
// 9. Проверить код через линтеры
// 10. Если сервер упал (сайт поменял домен, ошибки в веб-парсинге и т.д.), то сервис не должен перезапускаться
// 11. Настроить CI/CD
// 12. Написать тесты

func run() error {
	sMmrp := scrapers.MMRPScraper{}
	sMapm := scrapers.MAPMScraper{}
	cs, err := scrapers.Init()
	if err != nil {
		log.Fatal("Error in config", err)
	}
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(15).Minutes().Do(func() {
		sMmrp.Scrape(cs)

	})
	scheduler.Every(40).Minutes().Do(func() {
		sMapm.Scrape(cs)
	})
	scheduler.StartBlocking()
	return nil
}
