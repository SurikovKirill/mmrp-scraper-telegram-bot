package main

import (
	"fmt"
	"log"
	"mmrp-scraper/internal/scrapers"
	"mmrp-scraper/internal/telegram"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// TODO:
// 2. Создать эндпоинт, через который можно пинговать сервер и проверять работу
// (сделать так, чтобы при пинге выдавались логи)
// 3. Настроить graceful shutdown
// 5. Оптимизация докер-контейнера
// 8. Настроить CI/CD
// 9. Написать тесты
// 10. Настроить репозиторий на гитхабе
// 11. Отрефакторить в соотвествии с 12-факторной методологией

func run() error {
	// Инициализация парсеров
	sMmrp := scrapers.MMRPScraper{}
	sMapm := scrapers.MAPMScraper{}
	// Инициализация конфига mapm логопас
	if err := sMapm.Init(); err != nil {
		return err
	}
	// Инициализация конфига телеграма
	ct, err := telegram.Init()
	if err != nil {
		return err
	}
	// Создание кронов
	log.Println("Starting scheduler")
	fmt.Println(time.Local)
	scheduler := gocron.NewScheduler(time.Local)
	// Скраппинг MMRP каждые 15 минут
	scheduler.Every(15).Minutes().Do(func() {
		log.Println("Start MMRP task")
		sMmrp.Scrape(ct)
	})
	// Скраппинг MAPM каждые 10 часов
	scheduler.Cron("15 7,14 * * *").Do(func() {
		log.Println("Start MAPM task")
		sMapm.ScrapeWithRod(ct)
	})
	scheduler.StartBlocking()
	return nil
}
