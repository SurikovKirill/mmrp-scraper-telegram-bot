package main

import (
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

func run() error {
	sMmrp := scrapers.MMRPScraper{}
	sMapm := scrapers.MAPMScraper{}
	if err := sMapm.Init(); err != nil {
		return err
	}
	ct, err := telegram.Init()
	if err != nil {
		return err
	}
	log.Println("Starting scheduler")
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(15).Minutes().Do(func() {
		log.Println("Start MMRP task")
		sMmrp.Scrape(ct)
	})
	scheduler.Cron("15 10,17 * * *").Do(func() {
		log.Println("Start MAPM task")
		sMapm.ScrapeWithRod(ct)
	})
	scheduler.StartBlocking()
	return nil
}
