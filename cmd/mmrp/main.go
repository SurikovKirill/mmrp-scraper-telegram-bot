package main

import (
	"mmrp-scraper/internal/scrapers"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
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
	// Открыть файл для сбора логов
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)

	// Инициализация парсеров
	sMmrp := scrapers.MMRPScraper{}
	// sMapm := scrapers.MAPMScraper{}
	cs, err := scrapers.Init()
	if err != nil {
		log.WithFields(log.Fields{"package": "main", "function": "run", "error": err}).Error("Error in config")
		return err
	}
	f.Close()

	// Создание кронов
	scheduler := gocron.NewScheduler(time.UTC)

	// Скраппинг MMRP каждые 15 минут
	scheduler.Every(15).Minutes().Do(func() {
		log.New().Print("Start MMRP task")
		sMmrp.Scrape(cs)
	})

	// Скраппинг MAPM каждые 40 минут
	scheduler.Every(40).Minutes().Do(func() {
		log.New().Print("Start MAPM task")
		// sMapm.ScrapeTable(cs)
	})
	scheduler.StartBlocking()
	return nil
}
