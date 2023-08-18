package main

import (
	"log"
	"mmrp-scraper/internal/scrapers/mapm"
	"mmrp-scraper/internal/scrapers/mmrp"
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
// 1. Оптимизация докер-контейнера
// 2. Настроить CI/CD

func run() error {
	// Инициализация парсеров
	sMmrp := mmrp.Scraper{}
	sMapm := mapm.Scraper{}
	// Инициализация конфига mapm (Логопас)
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
	scheduler := gocron.NewScheduler(time.Local)
	//sMmrp.Scrape(ct)
	sMapm.ScrapeWithRod(ct)
	// Скраппинг MMRP каждые 15 минут
	scheduler.Every(15).Minutes().Do(func() {
		log.Println("Start MMRP task")
		sMmrp.Scrape(ct)
	})
	// Скраппинг MAPM каждые 10 часов
	scheduler.Cron("15 10,17 * * *").Do(func() {
		log.Println("Start MAPM task")
		//sMapm.ScrapeWithRod(ct)
	})
	scheduler.StartBlocking()
	return nil
}
